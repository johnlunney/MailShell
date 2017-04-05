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

# License
Copyright 2017 Nabeel Omer <nabeelkomer@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
