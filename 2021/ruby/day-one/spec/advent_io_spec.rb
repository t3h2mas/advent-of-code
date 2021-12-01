require 'advent_io'

RSpec.describe 'AdventIO' do
  describe 'as_lines' do
    let(:file) { StringIO.new('\n123\n321\n\n\n') }

    it 'reads lines while handling whitespace' do
      expect(AdventIO.as_lines(file).to eq(['123', '321'])
    end
  end
end
