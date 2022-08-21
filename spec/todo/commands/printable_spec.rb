require "spec_helper"
require "stringio"

RSpec.describe Todo::Commands::Printable do
  let(:printer_klass) do
    Class.new do
      include Todo::Commands::Printable

      def output
        @output ||= StringIO.new
      end
    end
  end

  let(:printer) { printer_klass.new }

  describe "#print_todos" do
    context "when todos have no subtodos" do
      let(:todos) do
        [
          Todo::Todo.new(id: 2, title: "dummy", state: :waiting),
          Todo::Todo.new(id: 1, title: "dummy", state: :undone),
          Todo::Todo.new(id: 10, title: "dummy", state: :done)
        ]
      end

      it "prints todos to output at the same level" do
        printer.print_todos(todos)
        expect(printer.output.string).to eq(<<~TEXT)
           2 | \e[2mdummy\e[0m
           1 | dummy
          10 | \e[2;9mdummy\e[0m
        TEXT
      end
    end

    context "when todos have subtodos" do
      let(:todos) do
        [
          Todo::Todo.new(id: 1, title: "dummy", subtodos: [
            Todo::Todo.new(id: 2, title: "dummy", subtodos: [
              Todo::Todo.new(id: 3, title: "dummy")
            ])
          ])
        ]
      end

      it "prints todos in nested form" do
        printer.print_todos(todos)
        expect(printer.output.string).to eq(<<~TEXT)
          1 | dummy
              2 | dummy
                  3 | dummy
        TEXT
      end
    end

    context "when todos have tags" do
      let(:todos) do
        [
          Todo::Todo.new(id: 1, title: "dummy", state: :undone, tags: ["dummy1"]),
          Todo::Todo.new(id: 2, title: "dummy", state: :done, tags: ["dummy1", "dummy2"]),
          Todo::Todo.new(id: 3, title: "dummy", state: :waiting, tags: ["dummy1", "dummy2"])
        ]
      end

      it "prints todos with tags" do
        printer.print_todos(todos)
        expect(printer.output.string).to eq(<<~TEXT)
          1 | dummy \e[36m#dummy1\e[0m
          2 | \e[2;9mdummy \e[36m#dummy1 \e[36m#dummy2\e[0m
          3 | \e[2mdummy \e[36m#dummy1 \e[36m#dummy2\e[0m
        TEXT
      end
    end
  end
end
