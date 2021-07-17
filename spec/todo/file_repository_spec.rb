require "spec_helper"

require "json"
require "pathname"
require "tmpdir"

RSpec.describe Todo::FileRepository do
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
          Todo::FileRepository.new(root_path: Pathname.pwd)
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

        Todo::FileRepository.new(root_path: Pathname.pwd)
        expect(index_path.read.chomp).to eq(index_json)
      end
    end

    context "when archived directory doesn't exist" do
      it "creates archived directory" do
        archived_path = Pathname.pwd.join("archived")
        expect {
          Todo::FileRepository.new(root_path: Pathname.pwd)
        }.to change { archived_path.exist? }.from(false).to(true)
        expect(archived_path).to be_directory
      end
    end

    context "when archived directory exists" do
      it "doesn't raise any errors" do
        archived_path = Pathname.pwd.join("archived")
        archived_path.mkdir

        expect {
          Todo::FileRepository.new(root_path: Pathname.pwd)
        }.not_to raise_error
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
      repository = Todo::FileRepository.new(root_path: Pathname.pwd)

      todo1_path = Pathname.pwd.join("1.md")
      expect { repository.create(title: "dummy 1") }.to change { todo1_path.exist? }.from(false).to(true)
      expect(todo1_path.read).to eq(<<~TEXT)
        ---
        title: dummy 1
        status: undone
        ---


      TEXT

      todo2_path = Pathname.pwd.join("2.md")
      expect { repository.create(title: "dummy 2") }.to change { todo2_path.exist? }.from(false).to(true)
    end

    it "updates index file" do
      repository = Todo::FileRepository.new(root_path: Pathname.pwd)
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
end
