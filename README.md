# MailShell v2.0
MailShell is a daemon that allows you to execute shell commands over your email.

It's useful in situations where remote access is required and using SSH is not possible, for example on machines with dynamic IP addresses.

Binaries for macOS and Linux are [available](https://github.com/nabeelomer/MailShell/releases/tag/v2.0).

# Usage
```zsh
$ mailshell -help
```

Send an email to the email account you used with MailShell with the subject line in the format `mailshell [shell command]` and MailShell will run the command and reply with the output.

MailShell checks the last 4 emails for that subject line.

# License
GPL v3

Copyright (C) 2017 Nabeel Omer
