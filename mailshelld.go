/*
	Copyright (C) 2017 Nabeel Omer

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	imap "github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func checkForCommandEmail(uname string, password string, server string) string {
	//I haven't yet figured out how to use go-imap, this was an example in the readme that I stole and used.
	c, err := client.DialTLS(server, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(uname, password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// We're using unsigned integers here, only substract if the result is > 0
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	var done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []string{imap.EnvelopeMsgAttr}, messages)
	}()

	for msg := range messages {
		if strings.Split(msg.Envelope.Subject, "")[0] == "mailshell" {
			return strings.Split(msg.Envelope.Subject, "|")[1]
		}
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
	return ""
}

func main() {
	var server = flag.String("imapserver", "imap.gmail.com:993", "IMAP server address, defaults to Gmail")
	var uname = flag.String("username", "", "Email Username")
	var password = flag.String("password", "", "Email Password")
	var shell = flag.String("shell", "/bin/sh", "Path to the shell you prefer, defaults to `sh`")
	// var frequency = flag.Int("pollFrequency", 2, "How frequently should mailshelld poll your email")
	flag.Parse()

	if *uname == "" || *password == "" {
		log.Fatal("Empty username and password")
		return
	}

	for {
		var command = checkForCommandEmail(*uname, *password, *server)
		if command != "" {
			log.Println("Executing command: " + command)
			func() {
				cmd := exec.Command(*shell)
				cmd.Stdin = strings.NewReader(command)
				var out bytes.Buffer
				cmd.Stdout = &out
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(out.String())
			}()
		} else {
			log.Println("Nothing to do")
		}
		log.Println("Going to sleep")
		time.Sleep(2 * time.Minute)
	}
}
