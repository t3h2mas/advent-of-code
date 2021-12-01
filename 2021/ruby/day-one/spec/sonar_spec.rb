require 'sonar'

RSpec.describe Sonar do
  describe '.depth_change_rate' do
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
      expect(described_class.depth_change_rate(measurements)).to eq(7)
    end

    it 'handles a negative change_rate' do
      expect(described_class.depth_change_rate(measurements.reverse)).to eq(-7)
    end
  end
end
