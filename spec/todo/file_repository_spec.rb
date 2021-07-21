require "spec_helper"

require "json"
require "pathname"
require "stringio"
require "tmpdir"

RSpec.describe Todo::FileRepository do
  let(:output) { StringIO.new }
  let(:error_output) { StringIO.new }

  describe "#initialize" do
    around do |example|
      Dir.mktmpdir do |dir|
        Dir.chdir(dir) do
          example.run
        end
      end
    end

    context "when index file doesn't exist" do
      it "creates index file" do
        index_path = Pathname.pwd.join("index.json")
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
      it "doesn't overwrite" do
        index_json = JSON.pretty_generate({
          todos: {"": [1]},
          archived: {},
          metadata: {
            lastId: 1,
            missingIds: []
          }
        })
        index_path = Pathname.pwd.join("index.json")
        index_path.open("wb") { |file| file.puts(index_json) }

        Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
        expect(index_path.read.chomp).to eq(index_json)
      end
    end

    context "when archived directory doesn't exist" do
      it "creates archived directory" do
        archived_path = Pathname.pwd.join("archived")
        expect {
          Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
        }.to change { archived_path.exist? }.from(false).to(true)
        expect(archived_path).to be_directory
      end
    end

    context "when archived directory exists" do
      it "doesn't raise any errors" do
        archived_path = Pathname.pwd.join("archived")
        archived_path.mkdir

        expect {
          Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
        }.not_to raise_error
      end
    end
  end

  describe "#list" do
    around do |example|
      Dir.mktmpdir do |dir|
        Dir.chdir(dir) do
          example.run
        end
      end
    end

    context "when todo file isn't broken" do
      it "returns todos" do
        repository = Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
        repository.create(title: "dummy 1")
        repository.create(title: "dummy 2")

        todos = repository.list
        expect(todos).to contain_exactly(
          an_instance_of(Todo::Todo).and(having_attributes(id: 1)),
          an_instance_of(Todo::Todo).and(having_attributes(id: 2))
        )
      end
    end

    context "when index file include ID but the todo file is not found" do
      it "puts warning message to error output" do
        repository = Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
        index_path = Pathname.pwd.join("index.json")
        index = JSON.parse(index_path.read, symbolize_names: true)
        index[:todos][:""] = [1]
        index[:metadata][:lastId] = 1
        index_json = JSON.pretty_generate(index)
        index_path.open("wb") { |file| file.puts(index_json) }

        repository.list
        todo_path = Pathname.pwd.join("1.md")
        expect(error_output.string).to eq("todo file is not found: #{todo_path}\n")
      end
    end

    context "when todo file is broken" do
      it "puts warning message to error output" do
        repository = Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
        repository.create(title: "dummy")
        todo_path = Pathname.pwd.join("1.md")
        todo_path.open("wb+") { |file| file.puts("") }

        repository.list
        expect(error_output.string).to eq("todo file is broken: #{todo_path}\n")
      end
    end

    context "when front matter of todo file is broken" do
      it "puts warning message to error output" do
        repository = Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
        repository.create(title: "dummy")
        todo_path = Pathname.pwd.join("1.md")
        todo_path.open("wb+") do |file|
          file.puts(<<~TEXT)
            ---
            title: %
            state: undone
            ---

            body
          TEXT
        end

        repository.list
        expect(error_output.string).to eq("todo file is broken: #{todo_path}\n")
      end
    end
  end

  describe "#create" do
    around do |example|
      Dir.mktmpdir do |dir|
        Dir.chdir(dir) do
          example.run
        end
      end
    end

    it "creates a file" do
      repository = Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)

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
      repository = Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
      index_path = Pathname.pwd.join("index.json")

      original_index = {todos: {}, archived: {}, metadata: {lastId: 0, missingIds: []}}
      expected_index1 = {todos: {"": [1]}, archived: {}, metadata: {lastId: 1, missingIds: []}}
      expect {
        repository.create(title: "dummy 1")
      }.to change {
        JSON.parse(index_path.read, symbolize_names: true)
      }.from(original_index).to(expected_index1)

      expected_index2 = {todos: {"": [1, 2]}, archived: {}, metadata: {lastId: 2, missingIds: []}}
      expect {
        repository.create(title: "dummy 2")
      }.to change {
        JSON.parse(index_path.read, symbolize_names: true)
      }.from(expected_index1).to(expected_index2)
    end
  end

  describe "#delete" do
    around do |example|
      Dir.mktmpdir do |dir|
        Dir.chdir(dir) do
          example.run
        end
      end
    end

    context "when todo file with given ID doesn't exist" do
      it "outputs message to error output" do
        repository = Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
        repository.delete(ids: [100])

        todo_path = Pathname.pwd.join("100.md")
        expect(error_output.string).to eq("todo file is not found: #{todo_path}\n")
      end
    end

    context "when todo file with given ID exists" do
      before do
        repository.create(title: "dummy")
      end

      let(:repository) do
        Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
      end

      it "deletes the todo file" do
        todo_path = Pathname.pwd.join("1.md")
        expect { repository.delete(ids: [1]) }.to change { todo_path.exist? }.from(true).to(false)
      end

      it "updates index file" do
        index_path = Pathname.pwd.join("index.json")
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
