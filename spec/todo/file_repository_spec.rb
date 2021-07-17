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
  end
end
