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
	"net/smtp"
	"os/exec"
	"strings"
	"time"

	imap "github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func checkForCommandEmail(uname string, password string, imaps string) string {
	//I haven't yet figured out how to use go-imap, this was an example in the readme that I stole and used.
	c, err := client.DialTLS(imaps, nil)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(uname, password); err != nil {
		log.Fatal(err)
		return ""
	}
	log.Println("Logged in")

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
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
		if strings.Split(msg.Envelope.Subject, " ")[0] == "mailshell" && (msg.Envelope.Sender[0].MailboxName+"@"+msg.Envelope.Sender[0].HostName) == uname {
			var array = strings.Split(msg.Envelope.Subject, " ")
			return strings.Join(append(array[:0], array[0+1:]...)[:], " ")
		}
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
	return ""
}

func respond(uname string, password string, smtps string, text string) {
	auth := smtp.PlainAuth(
		"",
		uname,
		password,
		strings.Split(smtps, ":")[0],
	)
	err := smtp.SendMail(
		smtps,
		auth,
		uname,
		[]string{uname},
		[]byte("To: "+uname+"\r\n"+"Subject: Mailshell Command Result\r\n"+"\r\n"+text+"\r\n"),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response Complete")
}

func main() {
	var smtps = flag.String("smtpServer", "smtp.gmail.com:587", "SMTP server address")
	var imaps = flag.String("imapServer", "imap.gmail.com:993", "IMAP server address")
	var uname = flag.String("username", "", "Email Username")
	var password = flag.String("password", "", "Email Password")
	var shell = flag.String("shell", "/bin/sh", "Path to the shell you prefer")
	var help = flag.Bool("help", false, "Show this help text")
	// var frequency = flag.Int("pollFrequency", 2, "How frequently should mailshelld poll your email")
	flag.Parse()

	if *help {
		fmt.Println("Usage: mailshell [options]\n\nOptions:")
		flag.PrintDefaults()
		return
	}

	if *uname == "" || *password == "" {
		log.Fatal("Empty username and password")
		return
	}

	log.Println("Mailshell v2.0\nShell: " + *shell + "\nServer: " + *imaps + "\nUsername: " + *uname)

	for {
		var command = checkForCommandEmail(*uname, *password, *imaps)
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
				respond(*uname, *password, *smtps, out.String())
			}()
		} else {
			log.Println("Nothing to do")
		}
		log.Println("Going to sleep")
		time.Sleep(2 * time.Minute)
	}
}
