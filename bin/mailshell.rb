# MIT License

# Copyright 2017 Nabeel Omer <nabeelkomer@gmail.com>
# Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
# documentation files (the "Software"), to deal in the Software without restriction, including without limitation
# the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and
# to permit persons to whom the Software is furnished to do so, subject to the following conditions:
# The above copyright notice and this permission notice shall be included in all copies or substantial portions
# of the Software.
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT
# LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO 
# EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN
# AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
# OTHER DEALINGS IN THE SOFTWARE.

#! /usr/bin/ruby
require 'gmail'

# Process.daemon

Signal.trap 'TERM' do
    exit 0
end

Signal.trap 'INT' do
    exit 0
end

if ARGV.length < 2
    puts 'Usage: mailshell.rb <email> <password>'
    exit -1
end

EMAIL = ARGV[0].freeze
PASSWORD = ARGV[1].freeze

ex = nil

gmail = Gmail.connect EMAIL, PASSWORD
loop do
    puts "loggedin? #{gmail.logged_in?}"
    begin
        gmail = Gmail.connect EMAIL, PASSWORD unless gmail.logged_in?
        puts "Connected!"
        # Look for a new command email
        gmail.inbox.emails :unread, from: EMAIL, gm: 'subject:mailshell' do |email|
            puts email.message.subject.to_s
            result = `#{email.message.subject.sub(/mailshell[ ]/, '')}`
            emailout = gmail.compose do
                to EMAIL
                subject "Result of #{email.message.subject.sub(/mailshell[ ]/, '')}"
                if !ex.nil?
                    body "#{result}\n\nLast Exception: #{ex}"
                    ex = nil
                else
                    body result.to_s
                end
            end
            emailout.deliver!
            email.read!
        end
        puts 'going to sleep'
        sleep 60
        puts 'woke up'
    rescue => e
        puts "rescue #{e}"
        ex = e
        sleep 60
        next
    end
end
gmail.logout
