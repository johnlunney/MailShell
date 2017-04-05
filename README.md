# MailShell
MailShell is a program that allows you to execute shell commands over your email.
You add your credentials to the `PASSWORD` and `EMAIL` fields at the top and execute the program.

MailShell currently only works with a Gmail account.

# Usage
After adding your credentials to the fields specified above and executing the program, send an email to yourself on your
Gmail account with the following syntax in the subject line:
```
mailshell <command>
```
MailShell will then parse the command, execute it, and then email you the output of the command.
It does not matter what the body of the email contains.

MailShell looks for emails starting with the string `mailshell` every two minutes.

MailShell will delete the command email after sending you the output.
