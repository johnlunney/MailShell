# MailShell
MailShell is a Linux daemon that allow you to execute shell commands over your email.

# Usage
```zsh
 ./mailshell --help
Usage: mailshell [options]

Options:
  -help
        Show this help text
  -imapServer string
        IMAP server address (default "imap.gmail.com:993")
  -password string
        Email Password
  -shell string
        Path to the shell you prefer (default "/bin/sh")
  -smtpServer string
        SMTP server address (default "smtp.gmail.com:587")
  -username string
        Email Username
```

```zsh
# Basic Usage:
mailshell -username="example@example.com" -password="password" -shell="zsh" -smtpServer="smtp.example.org" -imapServer="imap.example.org"
```

# License
GPL v3

Copyright (C) 2017 Nabeel Omer