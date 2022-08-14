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

    shared_context "when title includes square brackets" do
      let(:title) { "[dummy]" }
    end

    shared_context "when title includes colons" do
      let(:title) { ":dummy" }
    end

    shared_context "when title doesn't any special characters" do
      let(:title) { "dummy" }
    end

    shared_context "when tags are empty" do
      let(:tags) { [] }
    end

    shared_context "when tags are present" do
      let(:tags) { ["dummy"] }
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

    shared_context "when position is nil" do
      let(:position) { nil }
    end

    shared_context "when position is -1" do
      let(:position) { -1 }

      before do
        index_json = JSON.pretty_generate({
          todos: {"": [1]},
          archived: {},
          metadata: {
            lastId: 1,
            missingIds: []
          }
        })
        index_path.open("wb") { |file| file.puts(index_json) }
      end
    end

    shared_context "when position is 0" do
      let(:position) { 0 }

      before do
        index_json = JSON.pretty_generate({
          todos: {"": [1]},
          archived: {},
          metadata: {
            lastId: 1,
            missingIds: []
          }
        })
        index_path.open("wb") { |file| file.puts(index_json) }
      end
    end

    shared_context "when position is larger than or equal to the length of todos" do
      let(:position) { 10 }

      before do
        index_json = JSON.pretty_generate({
          todos: {"": [1]},
          archived: {},
          metadata: {
            lastId: 1,
            missingIds: []
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

    context "when missing IDs are empty and position is 0" do
      include_context "when title doesn't any special characters"
      include_context "when tags are empty"
      include_context "when missing IDs are empty"
      include_context "when position is 0"
      include_context "when parent_id isn't given"

      it "updates an index file" do
        expect {
          repository.create(title: title, tags: tags, position: position, parent_id: parent_id)
        }.to change {
          JSON.parse(index_path.read, symbolize_names: true)
        }.to({
          todos: {
            "": [2, 1]
          },
          archived: {},
          metadata: {
            lastId: 2,
            missingIds: []
          }
        })
      end
    end

    context "when missing IDs are empty and position is -1" do
      include_context "when title doesn't any special characters"
      include_context "when tags are empty"
      include_context "when missing IDs are empty"
      include_context "when position is -1"
      include_context "when parent_id isn't given"

      it "updates an index file" do
        expect {
          repository.create(title: title, tags: tags, position: position, parent_id: parent_id)
        }.to change {
          JSON.parse(index_path.read, symbolize_names: true)
        }.to({
          todos: {
            "": [1, 2]
          },
          archived: {},
          metadata: {
            lastId: 2,
            missingIds: []
          }
        })
      end
    end

    context "when missing IDs are empty and position is larger than or equal to the length of todos" do
      include_context "when title doesn't any special characters"
      include_context "when tags are empty"
      include_context "when missing IDs are empty"
      include_context "when position is larger than or equal to the length of todos"
      include_context "when parent_id isn't given"

      it "updates an index file" do
        expect {
          repository.create(title: title, tags: tags, position: position, parent_id: parent_id)
        }.to change {
          JSON.parse(index_path.read, symbolize_names: true)
        }.to({
          todos: {
            "": [1, 2]
          },
          archived: {},
          metadata: {
            lastId: 2,
            missingIds: []
          }
        })
      end
    end

    context "when missing IDs are empty and parent_id isn't given" do
      include_context "when title doesn't any special characters"
      include_context "when tags are empty"
      include_context "when missing IDs are empty"
      include_context "when position is nil"
      include_context "when parent_id isn't given"

      it "creates a todo file" do
        todo_path = Pathname.pwd.join("1.md")
        expect {
          repository.create(title: title, tags: tags, position: position, parent_id: parent_id)
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
          repository.create(title: title, position: position, parent_id: parent_id)
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

    context "when title includes square brackets, missing IDs are empty and parent_id isn't given" do
      include_context "when title includes square brackets"
      include_context "when tags are empty"
      include_context "when missing IDs are empty"
      include_context "when position is nil"
      include_context "when parent_id isn't given"

      it "creates a todo file with the title double-quoted" do
        todo_path = Pathname.pwd.join("1.md")
        expect {
          repository.create(title: title, tags: tags, position: position, parent_id: parent_id)
        }.to change { todo_path.exist? }.from(false).to(true)

        expect(todo_path.read).to eq(<<~TEXT)
          ---
          title: "[dummy]"
          state: undone
          ---


        TEXT
      end
    end

    context "when title includes square brackets, missing IDs are empty and parent_id isn't given" do
      include_context "when title includes colons"
      include_context "when tags are empty"
      include_context "when missing IDs are empty"
      include_context "when position is nil"
      include_context "when parent_id isn't given"

      it "creates a todo file with the title double-quoted" do
        todo_path = Pathname.pwd.join("1.md")
        expect {
          repository.create(title: title, tags: tags, position: position, parent_id: parent_id)
        }.to change { todo_path.exist? }.from(false).to(true)

        expect(todo_path.read).to eq(<<~TEXT)
          ---
          title: ":dummy"
          state: undone
          ---


        TEXT
      end
    end

    context "when tags are present, missing IDs are empty and parent_id isn't given" do
      include_context "when title doesn't any special characters"
      include_context "when tags are present"
      include_context "when missing IDs are empty"
      include_context "when position is nil"
      include_context "when parent_id isn't given"

      it "creates a todo file with tags" do
        todo_path = Pathname.pwd.join("1.md")
        expect {
          repository.create(title: title, tags: tags, position: position, parent_id: parent_id)
        }.to change { todo_path.exist? }.from(false).to(true)

        expect(todo_path.read).to eq(<<~TEXT)
          ---
          title: dummy
          state: undone
          tags: ["dummy"]
          ---


        TEXT
      end
    end

    context "when missing IDs are empty and parent_id is given" do
      include_context "when title doesn't any special characters"
      include_context "when tags are empty"
      include_context "when missing IDs are empty"
      include_context "when position is nil"
      include_context "when parent_id is given"

      it "creates a todo file" do
        todo_path = Pathname.pwd.join("2.md")
        expect {
          repository.create(title: title, tags: tags, position: position, parent_id: parent_id)
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
          repository.create(title: title, position: position, parent_id: parent_id)
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
      include_context "when title doesn't any special characters"
      include_context "when tags are empty"
      include_context "when missing IDs are present"
      include_context "when position is nil"
      include_context "when parent_id isn't given"

      it "creates a todo file with a missing ID" do
        todo_path = Pathname.pwd.join("#{missing_id}.md")
        expect {
          repository.create(title: title, tags: tags, position: position, parent_id: parent_id)
        }.to change { todo_path.exist? }.from(false).to(true)
      end

      it "removes a missing ID" do
        expect {
          repository.create(title: title, parent_id: parent_id)
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

    let!(:todo) { repository.create(title: "dummy") }
    let!(:subtodo) { repository.create(title: "dummy", parent_id: todo.id) }
    let!(:subsubtodo) { repository.create(title: "dummy", parent_id: subtodo.id) }

    context "when todo file with given ID doesn't exist" do
      let(:ids) { [100] }

      it "outputs message to error output" do
        repository.delete(ids: ids)

        todo_path = Pathname.pwd.join("#{ids.first}.md")
        expect(error_output.string).to eq("todo file is not found: #{todo_path}\n")
      end
    end

    context "when given ID is a parent todo's ID" do
      let(:ids) { [todo.id] }

      it "deletes all descendant todo files" do
        todo_path = Pathname.pwd.join("#{todo.id}.md")
        subtodo_path = Pathname.pwd.join("#{subtodo.id}.md")
        subsubtodo_path = Pathname.pwd.join("#{subsubtodo.id}.md")

        expect { repository.delete(ids: ids) }.to change { todo_path.exist? }.from(true).to(false)
          .and change { subtodo_path.exist? }.from(true).to(false)
          .and change { subsubtodo_path.exist? }.from(true).to(false)
      end

      it "updates index file" do
        before_index = {todos: {"": [todo.id], "#{todo.id}": [subtodo.id], "#{subtodo.id}": [subsubtodo.id]}, archived: {}, metadata: {lastId: subsubtodo.id, missingIds: []}}
        after_index = {todos: {"": []}, archived: {}, metadata: {lastId: subsubtodo.id, missingIds: [todo.id, subtodo.id, subsubtodo.id]}}

        expect {
          repository.delete(ids: ids)
        }.to change {
          JSON.parse(index_path.read, symbolize_names: true)
        }.from(before_index).to(after_index)
      end
    end

    context "when given IDs include a parent todo's ID and subtodo's ID" do
      let(:ids) { [todo.id, subtodo.id] }

      it "doesn't put any messages to error output" do
        repository.delete(ids: ids)
        expect(error_output.string).to be_empty
      end
    end
  end

  describe "#update" do
    let!(:repository) do
      Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
    end

    let!(:todo) { repository.create(title: "dummy") }
    let!(:subtodo) { repository.create(title: "dummy", parent_id: todo.id) }
    let!(:subsubtodo) { repository.create(title: "dummy", parent_id: subtodo.id) }

    context "when todo file with given ID doesn't exist" do
      let(:ids) { [100] }

      it "puts error message to error output" do
        repository.update(ids: ids, state: :undone)

        todo_path = Pathname.pwd.join("#{ids.first}.md")
        expect(error_output.string).to eq("todo file is not found: #{todo_path}\n")
      end
    end

    context "when given ID is a parent todo's ID" do
      let(:ids) { [todo.id] }

      it "updates all descendant todos" do
        todo_path = Pathname.pwd.join("#{todo.id}.md")
        subtodo_path = Pathname.pwd.join("#{subtodo.id}.md")
        subsubtodo_path = Pathname.pwd.join("#{subsubtodo.id}.md")

        expect { repository.update(ids: ids, state: :done) }.to change { todo_path.read.include?("state: done") }.to(true)
          .and change { subtodo_path.read.include?("state: done") }.to(true)
          .and change { subsubtodo_path.read.include?("state: done") }.to(true)
      end
    end

    context "when given IDs include a parent todo's ID and subtodo's ID" do
      let(:ids) { [todo.id, subtodo.id] }

      it "doesn't put any messages to error output" do
        repository.update(ids: ids, state: :done)
        expect(error_output.string).to be_empty
      end
    end
  end

  describe "#archive" do
    let!(:repository) do
      Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
    end

    let(:todo_id) { 1 }
    let(:todo_path) { Pathname.pwd.join("#{todo_id}.md") }
    let(:archived_todo_path) { archived_path.join("#{todo_id}.md") }
    let(:subtodo_id) { 2 }
    let(:subtodo_path) { Pathname.pwd.join("#{subtodo_id}.md") }
    let(:archived_subtodo_path) { archived_path.join("#{subtodo_id}.md") }
    let(:unknown_todo_id) { 100 }

    shared_context "when parent todo is undone" do
      before do
        todo_path.open("wb") do |file|
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

    shared_context "when parent todo is done" do
      before do
        todo_path.open("wb") do |file|
          file.puts(<<~TEXT)
            ---
            title: dummy
            state: done
            ---

            body
          TEXT
        end
      end
    end

    shared_context "when subtodo is undone" do
      before do
        subtodo_path.open("wb") do |file|
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

    shared_context "when subtodo is done" do
      before do
        subtodo_path.open("wb") do |file|
          file.puts(<<~TEXT)
            ---
            title: dummy
            state: done
            ---

            body
          TEXT
        end
      end
    end

    shared_context "when index file is normal" do
      before do
        index_json = JSON.pretty_generate({
          todos: {
            "": [todo_id],
            "#{todo_id}": [subtodo_id]
          },
          archived: {},
          metadata: {
            lastId: subtodo_id,
            missingIds: []
          }
        })
        index_path.open("wb") { |file| file.puts(index_json) }
      end
    end

    shared_context "when index file includes ID which todo file doesn't exist" do
      before do
        index_json = JSON.pretty_generate({
          todos: {
            "": [todo_id, unknown_todo_id],
            "#{todo_id}": [subtodo_id]
          },
          archived: {},
          metadata: {
            lastId: subtodo_id,
            missingIds: []
          }
        })
        index_path.open("wb") { |file| file.puts(index_json) }
      end
    end

    context "when both parent todo and subtodo is undone" do
      include_context "when parent todo is undone"
      include_context "when subtodo is undone"
      include_context "when index file is normal"

      it "doesn't move parent todo file and subtodo file" do
        expect { repository.archive }
          .to not_change { todo_path.exist? }
          .and not_change { subtodo_path.exist? }
      end

      it "doesn't update index file" do
        expect { repository.archive }.not_to change { JSON.parse(index_path.read, symbolize_names: true) }
      end
    end

    context "when parent todo is done but subtodo is undone" do
      include_context "when parent todo is done"
      include_context "when subtodo is undone"
      include_context "when index file is normal"

      it "moves both parent todo file and subtodo file into archived directory" do
        expect { repository.archive }
          .to change { todo_path.exist? }.to(false)
          .and change { archived_todo_path.exist? }.to(true)
          .and change { subtodo_path.exist? }.to(false)
          .and change { archived_subtodo_path.exist? }.to(true)
      end

      it "updates index file" do
        expect { repository.archive }.to change { JSON.parse(index_path.read, symbolize_names: true) }.to({
          todos: {},
          archived: {
            "": [todo_id],
            "#{todo_id}": [subtodo_id]
          },
          metadata: {
            lastId: subtodo_id,
            missingIds: []
          }
        })
      end
    end

    context "when parent todo is undone but subtodo is done" do
      include_context "when parent todo is undone"
      include_context "when subtodo is done"
      include_context "when index file is normal"

      it "moves only subtodo file into archived directory" do
        expect { repository.archive }
          .to not_change { todo_path.exist? }
          .and change { subtodo_path.exist? }.to(false)
          .and change { archived_subtodo_path.exist? }.to(true)
      end

      it "updates index file" do
        expect { repository.archive }.to change { JSON.parse(index_path.read, symbolize_names: true) }.to({
          todos: {
            "": [todo_id]
          },
          archived: {
            "#{todo_id}": [subtodo_id]
          },
          metadata: {
            lastId: subtodo_id,
            missingIds: []
          }
        })
      end
    end

    context "when both parent todo and subtodo is done" do
      include_context "when parent todo is done"
      include_context "when subtodo is done"
      include_context "when index file is normal"

      it "moves both parent todo file and subtodo file into archived directory" do
        expect { repository.archive }
          .to change { todo_path.exist? }.to(false)
          .and change { archived_todo_path.exist? }.to(true)
          .and change { subtodo_path.exist? }.to(false)
          .and change { archived_subtodo_path.exist? }.to(true)
      end

      it "updates index file" do
        expect { repository.archive }.to change { JSON.parse(index_path.read, symbolize_names: true) }.to({
          todos: {},
          archived: {
            "": [todo_id],
            "#{todo_id}": [subtodo_id]
          },
          metadata: {
            lastId: subtodo_id,
            missingIds: []
          }
        })
      end
    end

    context "when index file includes ID which todo file doesn't exist" do
      include_context "when parent todo is done"
      include_context "when subtodo is done"
      include_context "when index file includes ID which todo file doesn't exist"

      it "puts error message to error output" do
        repository.archive

        unknown_todo_path = Pathname.pwd.join("#{unknown_todo_id}.md")
        expect(error_output.string).to eq("todo file is not found: #{unknown_todo_path}\n")
      end
    end
  end

  describe "#move" do
    let!(:repository) do
      Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
    end

    before do
      index_path.open("wb") do |file|
        file.puts(JSON.pretty_generate({
          todos: {
            "": [1, 2, 3],
            "1": [4]
          },
          archived: {
            "": [5],
            "5": [6]
          },
          metadata: {
            lastId: 6,
            missingIds: []
          }
        }))
      end

      1.upto(6) do |id|
        todo_path = Pathname.pwd.join("#{id}.md")
        todo_path.open("wb") do |file|
          file.puts <<~TEXT
            ---
            title: dummy
            state: done
            ---

            body
          TEXT
        end
      end
    end

    [
      [[2, nil, -1], {todos: {"": [1, 3, 2], "1": [4]}}],
      [[2, nil, 1], {todos: {"": [2, 1, 3], "1": [4]}}],
      [[2, nil, 3], {todos: {"": [1, 3, 2], "1": [4]}}],
      [[2, nil, 4], {todos: {"": [1, 3, 2], "1": [4]}}],
      [[4, nil, 4], {todos: {"": [1, 2, 3, 4]}}],
      [[2, 1, 1], {todos: {"": [1, 3], "1": [2, 4]}}],
      [[2, 3, 1], {todos: {"": [1, 3], "1": [4], "3": [2]}}]
    ].each do |(id, parent_id, position), expected_index|
      context "when id: #{id}, parent_id: #{parent_id}, position: #{position}" do
        it "updates index file" do
          expect {
            repository.move(id: id, parent_id: parent_id, position: position)
          }.to change {
            JSON.parse(index_path.read, symbolize_names: true)
          }.to(a_hash_including(expected_index))
        end
      end
    end

    [
      [[100, nil, 1], "todo is not found: 100"],
      [[1, 100, 1], "parent is not found: 100"],
      [[1, 1, 1], "moving a todo under itself is forbidden"],
      [[5, nil, 1], "moving an archived todo is forbidden"],
      [[1, 6, 1], "moving a todo under an archived todo is forbidden"]
    ].each do |(id, parent_id, position), error_message|
      context "when id: #{id}, parent_id: #{parent_id}, position: #{position}" do
        it "puts error message to error output" do
          repository.move(id: id, parent_id: parent_id, position: position)
          expect(error_output.string).to eq(error_message + "\n")
        end
      end
    end
  end

  describe "#open" do
    let!(:repository) do
      Todo::FileRepository.new(root_path: Pathname.pwd, error_output: error_output)
    end

    context "when a todo with given ID exists" do
      let(:id) { 1 }
      let(:todo_path) { Pathname.pwd.join("#{id}.md") }

      before do
        FileUtils.touch(todo_path)
      end

      it "calls Kernel#.system with `open path/to/todo_file.md`" do
        expect(repository).to receive(:system).with("open #{todo_path}")
        repository.open(id: id)
      end
    end

    context "when a todo with given ID exists and is archived" do
      let(:id) { 1 }
      let(:todo_path) { archived_path.join("1.md") }

      before do
        FileUtils.touch(todo_path)
      end

      it "calls Kernel#.system with `open path/to/archived/todo_file.md`" do
        expect(repository).to receive(:system).with("open #{todo_path}")
        repository.open(id: id)
      end
    end

    context "when a todo with given ID doesn't exist" do
      let(:id) { 1 }

      it "puts error message to error output" do
        repository.open(id: id)
        expect(error_output.string).to eq("todo is not found: #{id}\n")
      end
    end
  end
end
