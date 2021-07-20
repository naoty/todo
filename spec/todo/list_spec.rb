require "spec_helper"

require "stringio"

RSpec.describe Todo::List do
  let(:output) { StringIO.new }
  let(:error_output) { StringIO.new }

  describe "#run" do
    context "when arguments are empty" do
      it "puts todos" do
        list = Todo::List.new(arguments: [], output: output, error_output: error_output)
        repository = instance_double(Todo::FileRepository)
        allow(repository).to receive(:list).and_return([
          Todo::Todo.new(id: 2, title: "dummy", state: :waiting, body: ""),
          Todo::Todo.new(id: 1, title: "dummy", state: :undone, body: ""),
          Todo::Todo.new(id: 10, title: "dummy", state: :done, body: "")
        ])

        list.run(repository: repository)
        expect(output.string).to eq(<<-TEXT)
   2 | \e[2mdummy\e[0m
   1 | dummy
  10 | \e[2;9mdummy\e[0m
        TEXT
      end
    end

    ["-h", "--help"].each do |flag|
      context "when arguments include #{flag} flag" do
        it "puts help message" do
          list = Todo::List.new(arguments: [flag], output: output, error_output: error_output)
          repository = instance_double(Todo::FileRepository)

          list.run(repository: repository)
          expect(output.string).to eq(Todo::List::HELP_MESSAGE)
        end
      end
    end
  end
end
