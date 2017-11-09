# MailShell v2.0
MailShell is a daemon that allows you to execute shell commands over your email.

Its useful in situations where remote access is required and using SSH is not possible. For example on machines with dynamic IP addresses.

# Usage
```zsh
$ mailshell -help
```

Send an email to the email account from itself with the subject line in the format `mailshell [shell command]` and mailshell will execute the shell command.

Mailshell checks the last 4 emails for that subject line.

# License
GPL v3

Copyright (C) 2017 Nabeel Omer