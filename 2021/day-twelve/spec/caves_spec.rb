# typed: false
# frozen_string_literal: true

require 'caves'

RSpec.describe Caves do
  let(:input) do
    <<~TEXT
      start-A
      start-b
      A-c
      A-b
      b-d
      A-end
      b-end
    TEXT
  end

  it 'finds the number of paths' do
    g = Caves.parse(input)
    expect(Explorer.paths(g)).to eq(10)
  end

  it 'parses the expected number of nodes' do
    g = Caves.parse(input)
    expect(g.nodes.count).to eq(6)
  end
end
