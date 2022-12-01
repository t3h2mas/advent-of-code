# frozen_string_literal: true

require 'consortium'

RSpec.describe Consortium do
  let(:input) do
    <<~TEXT
      5483143223
      2745854711
      5264556173
      6141336146
      6357385478
      4167524645
      2176841721
      6882881134
      4846848554
      5283751526
    TEXT
  end

  let(:small_input) do
    <<~TEXT
      11111
      19991
      19191
      19991
      11111
    TEXT
  end

  it 'counts flashes' do
    #expect(Consortium.new(small_input).cycle(2).flashes).to eq(9)
    #expect(Consortium.new(input).cycle(10).flashes).to eq(204)
    Consortium.new(input).cycle(195).flashes
  end
end
