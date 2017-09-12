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
