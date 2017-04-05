#! /usr/bin/ruby

require 'gmail'

exit unless File.new("/tmp/mailshell.lock").flock( File::LOCK_NB | File::LOCK_EX )

EMAIL = 'youremail@gmail.com'.freeze
PASSWORD = 'yourpassword'.freeze

puts "daemon pid: #{Process.pid}"
Process.daemon

ex = nil
gmail = Gmail.connect(EMAIL, PASSWORD)
loop do
    begin
        gmail = Gmail.connect(EMAIL, PASSWORD) unless gmail.logged_in?
        # Look for a new command email
        gmail.inbox.emails(:unread, from: EMAIL, gm: 'subject:mailshell') do |email|
            result = `#{email.message.subject.sub(/mailshell[ ]/, '')}`
            emailout = gmail.compose do
                to EMAIL
                subject "Result of #{email.message.subject.sub(/mailshell[ ]/, '')} — Mailshell"
                if ex != nil
                  body  "#{result.to_s}\n\nLast Exception: #{ex}"
                  ex = nil
                else
                  body "#{result.to_s}"
                end
            end
            emailout.deliver!
            email.delete!
        end
        # sleep for 2 mins
        sleep 60 * 2
    rescue => e
      ex = e
      sleep 60 * 2
      next
    end
end
gmail.logout
