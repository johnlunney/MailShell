# MailShell v2.0
MailShell is a daemon that allows you to execute shell commands over your email.

Its useful in situations where remote access is required and using SSH is not possible. For example on machines with dynamic IP addresses.

Binaries for macOS and Linux are [available](https://github.com/nabeelomer/MailShell/releases/tag/v2.0).

**Mailshell doesn't verify which account the command email came from.**

# Usage
```zsh
$ mailshell -help
```

Send an email to the email account you used with mailshell with the subject line in the format `mailshell [shell command]` and Mailshell will run the command and reply with the output.

Mailshell checks the last 4 emails for that subject line.

# License
GPL v3

Copyright (C) 2017 Nabeel Omer