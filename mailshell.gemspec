Gem::Specification.new do |s|
  s.name = 'MailShell'
  s.version     = '1.0.0'
  s.date        = Time.now.strftime('%Y-%m-%d').to_s
  s.summary     = 'Execute shell commands through your email'
  s.description = 'Execute shell commands through your email'
  s.authors     = ['Nabeel Omer']
  s.email       = 'nabeelkomer@gmail.com'
  s.platform    = Gem::Platform::RUBY
  s.files       = ['bin/mailshell.rb', 'MailShell.service', 'README.md']
  s.homepage    = 'https://github.com/nabeelomer/MailShell'
  s.license     = 'MIT'
  s.executables << 'mailshell.rb'
  s.add_runtime_dependency 'gmail'
  s.add_runtime_dependency 'mail'
  s.post_install_message = `cp MailShell.service /etc/systemd/system`
end
