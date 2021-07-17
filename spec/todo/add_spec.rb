require "spec_helper"
require "stringio"

RSpec.describe Todo::Add do
  describe "#run" do
    let(:output) { StringIO.new }
    let(:error_output) { StringIO.new }

    context "when arguments are empty" do
      it "puts help message to error output" do
        add = Todo::Add.new(arguments: [], output: output, error_output: error_output)
        add.run
      rescue SystemExit
        # ignore exit
      ensure
        expect(error_output.string).to eq(Todo::Add::HELP_MESSAGE)
      end

      it "exits with status code 1" do
        add = Todo::Add.new(arguments: [], output: output, error_output: error_output)
        expect { add.run }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
      end
    end

    ["-h", "--help"].each do |flag|
      context "when arguments include '#{flag}' flag" do
        it "puts help message to output" do
          add = Todo::Add.new(arguments: [flag], output: output, error_output: error_output)
          add.run
          expect(output.string).to eq(Todo::Add::HELP_MESSAGE)
        end
      end
    end
  end
end
