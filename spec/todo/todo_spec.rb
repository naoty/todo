require "spec_helper"

RSpec.describe Todo::Todo do
  describe "#should_be_archived?" do
    subject { todo.should_be_archived? }

    let(:parent) do
      Todo::Todo.new(
        id: 1,
        title: "dummy",
        state: parent_state,
        body: "body",
        subtodos: []
      )
    end

    let(:todo) do
      Todo::Todo.new(
        id: 2,
        title: "dummy",
        state: state,
        body: "body",
        subtodos: []
      )
    end

    let(:parent_state) { :undone }
    let(:state) { :undone }

    shared_context "when parent is nil" do
      before do
        parent.subtodos = []
        todo.parent = nil
      end
    end

    shared_context "when parent is present" do
      before do
        parent.subtodos << todo
        todo.parent = parent
      end
    end

    shared_context "when parent is undone" do
      let!(:parent_state) { :undone }
    end

    shared_context "when parent is done" do
      let!(:parent_state) { :done }
    end

    shared_context "when todo is undone" do
      let!(:state) { :undone }
    end

    shared_context "when todo is done" do
      let!(:state) { :done }
    end

    context "when parent is nil and todo is undone" do
      include_context "when parent is nil"
      include_context "when todo is undone"

      it { is_expected.to be false }
    end

    context "when parent is nil and todo is done" do
      include_context "when parent is nil"
      include_context "when todo is done"

      it { is_expected.to be true }
    end

    context "when parent is undone and todo is undone" do
      include_context "when parent is present"
      include_context "when parent is undone"
      include_context "when todo is undone"

      it { is_expected.to be false }
    end

    context "when parent is done and todo is undone" do
      include_context "when parent is present"
      include_context "when parent is done"
      include_context "when todo is undone"

      it { is_expected.to be true }
    end

    context "when parent is undone and todo is done" do
      include_context "when parent is present"
      include_context "when parent is undone"
      include_context "when todo is done"

      it { is_expected.to be true }
    end

    context "when parent is done and todo is done" do
      include_context "when parent is present"
      include_context "when parent is done"
      include_context "when todo is done"

      it { is_expected.to be true }
    end
  end
end
