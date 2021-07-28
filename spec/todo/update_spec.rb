require "spec_helper"
require "stringio"

RSpec.describe Todo::Update do
  describe "#run" do
    let(:output) { StringIO.new }
    let(:error_output) { StringIO.new }
    let(:repository) { instance_double(Todo::FileRepository) }

    context "when arguments are empty" do
      it "puts help message to error output" do
        done = Todo::Update.new(arguments: [], state: :done, output: output, error_output: error_output)
        done.run(repository: repository)
      rescue SystemExit
        # ignore exit
      ensure
        expect(error_output.string).to eq(done.help_message)
      end

      it "exits with status code 1" do
        done = Todo::Update.new(arguments: [], state: :done, output: output, error_output: error_output)
        expect {
          done.run(repository: repository)
        }.to raise_error(an_instance_of(SystemExit).and(having_attributes({status: 1})))
      end
    end

    ["-h", "--help"].each do |flag|
      context "when arguments include '#{flag}' flag" do
        it "puts help message to output" do
          done = Todo::Update.new(arguments: [flag], state: :done, output: output, error_output: error_output)
          done.run(repository: repository)
          expect(output.string).to eq(done.help_message)
        end
      end
    end

    context "when arguments include IDs" do
      it "calls Todo::FileRepository#update" do
        done = Todo::Update.new(arguments: ["1"], state: :done, output: output, error_output: error_output)
        expect(repository).to receive(:update).with(ids: [1], state: :done)
        done.run(repository: repository)
      end
    end
  end
end
