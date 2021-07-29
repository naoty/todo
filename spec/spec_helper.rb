require "bundler/setup"
Bundler.require(:test)

SimpleCov.start do
  enable_coverage :branch
  primary_coverage :branch

  add_filter "/spec/"
end

require "todo"

RSpec.configure do |config|
  config.example_status_persistence_file_path = "./spec/examples.txt"
end

RSpec::Matchers.define_negated_matcher :not_change, :change
