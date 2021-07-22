require "spec_helper"

require "fileutils"
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

    let(:todo_id) { 1 }
    let(:todo_path) { Pathname.pwd.join("#{todo_id}.md") }
    let(:subtodo_id) { 2 }
    let(:subtodo_path) { Pathname.pwd.join("#{subtodo_id}.md") }
    let(:not_found_todo_id) { 100 }

    shared_context "when a todo file exists" do
      before do
        FileUtils.touch(todo_path)
      end
    end

    shared_context "when a todo file is normal" do
      before do
        todo_path.open("wb+") do |file|
          file.puts(<<~TEXT)
            ---
            title: dummy
            state: undone
            ---

            body
          TEXT
        end
      end
    end

    shared_context "when a todo file is empty" do
      before do
        todo_path.truncate(0)
      end
    end

    shared_context "when a todo file doesn't include front matter" do
      before do
        todo_path.open("wb+") { |file| file.puts("body") }
      end
    end

    shared_context "when a todo file includes broken front matter" do
      before do
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
    end

    shared_context "when an index file is normal" do
      before do
        index_path.open("wb+") do |file|
          file.puts(JSON.pretty_generate({
            todos: {
              "": [todo_id]
            },
            archived: {},
            metadata: {
              lastId: todo_id,
              missingIds: []
            }
          }))
        end
      end
    end

    shared_context "when an index file is empty" do
      before do
        index_path.truncate(0)
      end
    end

    shared_context "when an index file includes ID which todo file doesn't exist" do
      before do
        index_path.open("wb+") do |file|
          file.puts(JSON.pretty_generate({
            todos: {
              "": [not_found_todo_id]
            },
            archived: {},
            metadata: {
              lastId: not_found_todo_id,
              missingIds: []
            }
          }))
        end
      end
    end

    shared_context "when a todo has a subtodo" do
      before do
        subtodo_path.open("wb+") do |file|
          file.puts(<<~TEXT)
            ---
            title: dummy
            state: undone
            ---

            body
          TEXT
        end

        index_path.open("wb+") do |file|
          file.puts(JSON.pretty_generate({
            todos: {
              "": [todo_id],
              "#{todo_id}": [subtodo_id]
            },
            archived: {},
            metadata: {
              lastId: subtodo_id,
              missingIds: []
            }
          }))
        end
      end
    end

    context "when todo doesn't exist" do
      it "returns empty array" do
        expect(repository.list).to be_empty
      end
    end

    context "when a todo exists but the todo file is empty" do
      include_context "when a todo file exists"
      include_context "when a todo file is empty"
      include_context "when an index file is normal"

      it "returns empty array" do
        expect(repository.list).to be_empty
      end

      it "puts message to error output" do
        repository.list
        expect(error_output.string).to eq("todo file is broken: #{todo_path}\n")
      end
    end

    context "when a todo exists but the todo file doesn't include front matter" do
      include_context "when a todo file exists"
      include_context "when a todo file doesn't include front matter"
      include_context "when an index file is normal"

      it "returns empty array" do
        expect(repository.list).to be_empty
      end

      it "puts message to error output" do
        repository.list
        expect(error_output.string).to eq("todo file is broken: #{todo_path}\n")
      end
    end

    context "when a todo exists but the todo file includes broken front matter" do
      include_context "when a todo file exists"
      include_context "when a todo file includes broken front matter"
      include_context "when an index file is normal"

      it "returns empty array" do
        expect(repository.list).to be_empty
      end

      it "puts message to error output" do
        repository.list
        expect(error_output.string).to eq("todo file is broken: #{todo_path}\n")
      end
    end

    context "when a todo exists but an index file is empty" do
      include_context "when a todo file exists"
      include_context "when a todo file is normal"
      include_context "when an index file is empty"

      it "returns empty array" do
        expect(repository.list).to be_empty
      end

      it "puts message to error output" do
        repository.list
        expect(error_output.string).to eq("index file is broken: #{index_path}\n")
      end
    end

    context "when a todo exists but an index file includes ID which todo file doesn't exist" do
      include_context "when a todo file exists"
      include_context "when a todo file is normal"
      include_context "when an index file includes ID which todo file doesn't exist"

      it "returns empty array" do
        expect(repository.list).to be_empty
      end

      it "puts message to error output" do
        repository.list
        not_found_todo_path = Pathname.pwd.join("#{not_found_todo_id}.md")
        expect(error_output.string).to eq("todo file is not found: #{not_found_todo_path}\n")
      end
    end

    context "when a todo file and an index file is normal" do
      include_context "when a todo file exists"
      include_context "when a todo file is normal"
      include_context "when an index file is normal"

      it "returns array containing a todo" do
        expect(repository.list).to contain_exactly(
          an_instance_of(Todo::Todo).and(having_attributes(
            id: todo_id
          ))
        )
      end
    end

    context "when a todo has a subtodo" do
      include_context "when a todo file exists"
      include_context "when a todo file is normal"
      include_context "when an index file is normal"
      include_context "when a todo has a subtodo"

      it "returns array containing a todo with a subtodo" do
        expect(repository.list).to contain_exactly(
          an_instance_of(Todo::Todo).and(having_attributes(
            id: todo_id,
            subtodos: a_collection_containing_exactly(
              an_instance_of(Todo::Todo).and(having_attributes(
                id: subtodo_id
              ))
            )
          ))
        )
      end
    end
  end

  describe "#create" do
    let!(:repository) do
      Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
    end

    shared_context "when missing IDs are empty" do
      before do
        index_json = JSON.pretty_generate({
          todos: {},
          archived: {},
          metadata: {
            lastId: 0,
            missingIds: []
          }
        })
        index_path.open("wb") { |file| file.puts(index_json) }
      end
    end

    shared_context "when missing IDs are present" do
      let(:missing_id) { 1 }

      before do
        index_json = JSON.pretty_generate({
          todos: {},
          archived: {},
          metadata: {
            lastId: 2,
            missingIds: [missing_id]
          }
        })
        index_path.open("wb") { |file| file.puts(index_json) }
      end
    end

    shared_context "when parent_id is given" do
      let!(:parent_id) do
        parent = repository.create(title: "dummy")
        parent.id
      end
    end

    shared_context "when parent_id isn't given" do
      let(:parent_id) { nil }
    end

    context "when missing IDs are empty and parent_id isn't given" do
      include_context "when missing IDs are empty"
      include_context "when parent_id isn't given"

      it "creates a todo file" do
        todo_path = Pathname.pwd.join("1.md")
        expect {
          repository.create(title: "dummy", parent_id: parent_id)
        }.to change { todo_path.exist? }.from(false).to(true)

        expect(todo_path.read).to eq(<<~TEXT)
          ---
          title: dummy
          state: undone
          ---


        TEXT
      end

      it "updates an index file" do
        expect {
          repository.create(title: "dummy", parent_id: parent_id)
        }.to change {
          JSON.parse(index_path.read, symbolize_names: true)
        }.to({
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
    end

    context "when missing IDs are empty and parent_id is given" do
      include_context "when missing IDs are empty"
      include_context "when parent_id is given"

      it "creates a todo file" do
        todo_path = Pathname.pwd.join("2.md")
        expect {
          repository.create(title: "dummy", parent_id: parent_id)
        }.to change { todo_path.exist? }.from(false).to(true)

        expect(todo_path.read).to eq(<<~TEXT)
          ---
          title: dummy
          state: undone
          ---


        TEXT
      end

      it "updates an index file" do
        expect {
          repository.create(title: "dummy", parent_id: parent_id)
        }.to change {
          JSON.parse(index_path.read, symbolize_names: true)
        }.to({
          todos: {
            "": [1],
            "1": [2]
          },
          archived: {},
          metadata: {
            lastId: 2,
            missingIds: []
          }
        })
      end
    end

    context "when missing IDs are present and parent_id isn't given" do
      include_context "when missing IDs are present"
      include_context "when parent_id isn't given"

      it "creates a todo file with a missing ID" do
        todo_path = Pathname.pwd.join("#{missing_id}.md")
        expect {
          repository.create(title: "dummy", parent_id: parent_id)
        }.to change { todo_path.exist? }.from(false).to(true)
      end

      it "removes a missing ID" do
        expect {
          repository.create(title: "dummy", parent_id: parent_id)
        }.to change {
          JSON.parse(index_path.read, symbolize_names: true)
        }.to({
          todos: {
            "": [missing_id]
          },
          archived: {},
          metadata: {
            lastId: 2,
            missingIds: []
          }
        })
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
