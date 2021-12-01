require 'sonar'

RSpec.describe Sonar do
  describe '.depth_increases' do
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

    it 'gets the depth change rate' do
      expect(described_class.depth_increases(measurements)).to eq(7)
    end
  end
end
