require 'sonar'

RSpec.describe Sonar do
  let(:measurements) do
    [
      199,
      200,
      208,
      210,
      200,
      207,
      240,
      269,
      260,
      263
    ]
  end

  describe '.depth_increases' do
    it 'gets the depth change rate' do
      expect(described_class.depth_increases(measurements)).to eq(7)
    end
  end

  describe '.depth_increases_over_window' do
    it 'gets the depth change rate' do
      expect(described_class.depth_increases_over_window(measurements)).to eq(5)
    end
  end
end
