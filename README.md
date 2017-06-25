# MailShell
MailShell is a Linux (systemd compatible) daemon that allow you to execute shell commands over your email.

MailShell depends on [ruby]() and the [gmail gem]().

# Installing MailShell
Installing MailShell is as simple as
```zsh
curl -c https://github.com/nabeelomer/MailShell/blob/master/MailShell-1.0.0.gem
gem install ./MailShell-1.0.0.gem
```
This will cause the service to be installed, configured, and start running.

Edit `/etc/mailshell` and add your Gmail username and password in the format `username@gmail.com|password` then restart
MailShell by running
```zsh
sudo service mailshell restart
```

# Using MailShell
To make MailShell execute a command send and email to your own Gmail account with the subject line written in the syntax
```zsh
mailshell <command>
```
This will cause MailShell to execute the <command> and email you back with the output of the command.
The contents of the body of the email do not matter.

MailShell will poll the Gmail account every two minutes to check for command emails.

# License
MIT License

Copyright 2017 Nabeel Omer <nabeelkomer@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
