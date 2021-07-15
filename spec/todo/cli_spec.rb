require "spec_helper"
require "stringio"

RSpec.describe Todo::CLI do
  describe "#run" do
    ["-h", "--help"].each do |flag|
      context "when arguments include '#{flag}' flag" do
        it "puts help message to output" do
          output = StringIO.new
          error_output = StringIO.new
          cli = Todo::CLI.new(
            arguments: [flag],
            output: output,
            error_output: error_output
          )
          cli.run
          expect(output.string).to eq(Todo::CLI::HELP_MESSAGE)
        end
      end
    end
  end
end
