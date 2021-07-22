require "spec_helper"

require "json"
require "pathname"
require "stringio"
require "tmpdir"

RSpec.describe Todo::FileRepository do
  let(:output) { StringIO.new }
  let(:error_output) { StringIO.new }
  let(:archived_path) { Pathname.pwd.join("archived") }
  let(:index_path) { Pathname.pwd.join("index.json") }

  around do |example|
    Dir.mktmpdir do |dir|
      Dir.chdir(dir) do
        example.run
      end
    end
  end

  describe "#initialize" do
    context "when index file doesn't exist" do
      it "creates index file" do
        expect {
          Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
        }.to change { index_path.exist? }.from(false).to(true)

        index = JSON.parse(index_path.read, symbolize_names: true)
        expect(index).to include({
          todos: {},
          archived: {},
          metadata: {
            lastId: 0,
            missingIds: []
          }
        })
      end
    end

    context "when index file exists" do
      let(:index_json) do
        JSON.pretty_generate({
          todos: {
            "": [1]
          },
          archived: {},
          metadata: {
            lastId: 1,
            missingIds: []
          }
        })
      end

      before do
        index_path.open("wb") { |file| file.puts(index_json) }
      end

      it "doesn't overwrite" do
        Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
        expect(index_path.read.chomp).to eq(index_json)
      end
    end

    context "when archived directory doesn't exist" do
      it "creates archived directory" do
        expect {
          Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
        }.to change { archived_path.exist? }.from(false).to(true)
        expect(archived_path).to be_directory
      end
    end

    context "when archived directory exists" do
      before do
        archived_path.mkdir
      end

      it "doesn't raise any errors" do
        expect {
          Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
        }.not_to raise_error
      end
    end
  end

  describe "#list" do
    let!(:repository) do
      Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
    end

    context "when todo file isn't broken" do
      context "when no todos have subtudos" do
        before do
          repository.create(title: "dummy 1")
          repository.create(title: "dummy 2")
        end

        it "returns only todos" do
          todos = repository.list
          expect(todos).to contain_exactly(
            an_instance_of(Todo::Todo).and(having_attributes(id: 1)),
            an_instance_of(Todo::Todo).and(having_attributes(id: 2))
          )
        end
      end

      context "when a todo have subtudos" do
        before do
          repository.create(title: "dummy 1")

          # TODO: Use #create to create subtudos
          repository.create(title: "dummy 2")
          index = JSON.parse(index_path.read, symbolize_names: true)
          index[:todos][:""] = [1]
          index[:todos][:"1"] = [2]
          index_json = JSON.pretty_generate(index)
          index_path.open("wb") { |file| file.puts(index_json) }
        end

        it "returns todos with subtodos" do
          todos = repository.list
          expect(todos).to contain_exactly(
            an_instance_of(Todo::Todo).and(having_attributes(
              id: 1,
              subtodos: a_collection_containing_exactly(
                an_instance_of(Todo::Todo).and(having_attributes(id: 2))
              )
            ))
          )
        end
      end
    end

    context "when index file include ID but the todo file is not found" do
      before do
        index = JSON.parse(index_path.read, symbolize_names: true)
        index[:todos][:""] = [1]
        index[:metadata][:lastId] = 1
        index_json = JSON.pretty_generate(index)
        index_path.open("wb") { |file| file.puts(index_json) }
      end

      it "puts warning message to error output" do
        repository.list
        todo_path = Pathname.pwd.join("1.md")
        expect(error_output.string).to eq("todo file is not found: #{todo_path}\n")
      end
    end

    context "when todo file is broken" do
      let(:todo_path) do
        Pathname.pwd.join("1.md")
      end

      before do
        repository.create(title: "dummy")
        todo_path.open("wb+") { |file| file.puts("") }
      end

      it "puts warning message to error output" do
        repository.list
        expect(error_output.string).to eq("todo file is broken: #{todo_path}\n")
      end
    end

    context "when front matter of todo file is broken" do
      let(:todo_path) do
        Pathname.pwd.join("1.md")
      end

      before do
        repository.create(title: "dummy")
        todo_path.open("wb+") do |file|
          file.puts(<<~TEXT)
            ---
            title: %
            state: undone
            ---

            body
          TEXT
        end
      end

      it "puts warning message to error output" do
        repository.list
        expect(error_output.string).to eq("todo file is broken: #{todo_path}\n")
      end
    end
  end

  describe "#create" do
    let!(:repository) do
      Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
    end

    context "when missing IDs are empty" do
      it "creates a file" do
        todo1_path = Pathname.pwd.join("1.md")
        expect { repository.create(title: "dummy 1") }.to change { todo1_path.exist? }.from(false).to(true)
        expect(todo1_path.read).to eq(<<~TEXT)
          ---
          title: dummy 1
          state: undone
          ---


        TEXT

        todo2_path = Pathname.pwd.join("2.md")
        expect { repository.create(title: "dummy 2") }.to change { todo2_path.exist? }.from(false).to(true)
      end

      it "updates index file" do
        original_index = {todos: {}, archived: {}, metadata: {lastId: 0, missingIds: []}}
        expected_index1 = {todos: {"": [1]}, archived: {}, metadata: {lastId: 1, missingIds: []}}
        expected_index2 = {todos: {"": [1, 2]}, archived: {}, metadata: {lastId: 2, missingIds: []}}

        expect {
          repository.create(title: "dummy 1")
        }.to change {
          JSON.parse(index_path.read, symbolize_names: true)
        }.from(original_index).to(expected_index1)

        expect {
          repository.create(title: "dummy 2")
        }.to change {
          JSON.parse(index_path.read, symbolize_names: true)
        }.from(expected_index1).to(expected_index2)
      end
    end

    context "when a missing ID exists" do
      let(:index) do
        {todos: {"": [2]}, archived: {}, metadata: {lastId: 2, missingIds: [1]}}
      end

      before do
        index_json = JSON.pretty_generate(index)
        index_path.open("wb") { |file| file.puts(index_json) }
      end

      it "creates todo file with the ID" do
        todo_path = Pathname.pwd.join("1.md")
        expect {
          repository.create(title: "dummy")
        }.to change { todo_path.exist? }.from(false).to(true)
      end

      it "removes the ID from missingIds" do
        expected_index = {todos: {"": [2, 1]}, archived: {}, metadata: {lastId: 2, missingIds: []}}
        expect {
          repository.create(title: "dummy")
        }.to change { JSON.parse(index_path.read, symbolize_names: true) }.from(index).to(expected_index)
      end
    end
  end

  describe "#delete" do
    let!(:repository) do
      Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
    end

    context "when todo file with given ID doesn't exist" do
      let(:ids) { [100] }

      it "outputs message to error output" do
        repository.delete(ids: ids)

        todo_path = Pathname.pwd.join("#{ids.first}.md")
        expect(error_output.string).to eq("todo file is not found: #{todo_path}\n")
      end
    end

    context "when todo file with given ID exists" do
      before do
        repository.create(title: "dummy")
      end

      it "deletes the todo file" do
        todo_path = Pathname.pwd.join("1.md")
        expect { repository.delete(ids: [1]) }.to change { todo_path.exist? }.from(true).to(false)
      end

      it "updates index file" do
        before_index = {todos: {"": [1]}, archived: {}, metadata: {lastId: 1, missingIds: []}}
        after_index = {todos: {"": []}, archived: {}, metadata: {lastId: 1, missingIds: [1]}}

        expect {
          repository.delete(ids: [1])
        }.to change {
          JSON.parse(index_path.read, symbolize_names: true)
        }.from(before_index).to(after_index)
      end
    end
  end
end
