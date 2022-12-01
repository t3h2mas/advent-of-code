require_relative '../lib/segments'

RSpec.describe Segments do
  describe '.parse' do
    let(:input) do
      <<~HEREDOC
    0,9 -> 5,9
    8,0 -> 0,8
    9,4 -> 3,4
    2,2 -> 2,1
    7,0 -> 7,4
    6,4 -> 2,0
    0,9 -> 2,9
    3,4 -> 1,4
    0,0 -> 8,8
    5,5 -> 8,2
      HEREDOC
    end
    context 'without diagonals' do
      let(:expected) do
        [[[0, 9], [5, 9]],
         [[9, 4], [3, 4]],
         [[2, 2], [2, 1]],
         [[7, 0], [7, 4]],
         [[0, 9], [2, 9]],
         [[3, 4], [1, 4]]].map(&:sort)
      end
      it 'returns the expected segments' do
        expect(Segments.parse(input, Segments.method(:validate_without_diagonals))).to eq(expected)
      end
    end

    context 'with diagonals' do
      it 'does'
    end
  end
  describe 'diagonal?' do
    it 'does' do
      expect(Segments.diagonal?([[1, 1], [3, 3]])).to eq(true)

      expect(Segments.diagonal?([[9, 7], [7, 9]])).to eq(true)
      expect(Segments.diagonal?([[1, 7], [7, 9]])).to eq(false)
    end
  end
end

