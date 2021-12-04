require 'bits'

RSpec.describe Bits do
  let(:numbies) do
    [
      "00100",
      "11110",
      "10110",
      "10111",
      "10101",
      "01111",
      "00111",
      "11100",
      "10000",
      "11001",
      "00010",
      "01010"
    ]
  end

  it 'computes the gamma' do
    expect(Bits.gamma(numbies)).to eq(22)
  end

  it 'computes the epsilon' do
    expect(Bits.epsilon(22)).to eq(9)
  end

  it 'finds the numbers' do
    expect(Bits.oxygen(numbies)).to eq(23)
    expect(Bits.scrubber(numbies)).to eq(10)
  end
end
