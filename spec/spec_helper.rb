require "bundler/setup"
Bundler.require(:test)

SimpleCov.start do
  enable_coverage :branch
  primary_coverage :branch
end

require "todo"

RSpec.configure do |config|
end
