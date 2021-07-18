require "bundler/setup"
Bundler.require(:test)

SimpleCov.start do
  enable_coverage :branch
  primary_coverage :branch
end

require "todo"

RSpec.configure do |config|
  config.example_status_persistence_file_path = "./spec/examples.txt"
end
