require 'chunk_validator'

RSpec.describe ChunkValidator do
  it 'detects incomplete chunks' do
    expect(described_class.validate('[({(<(())[]>[[{[]{<()<>>')[:result]).to eq(:incomplete)
  end

  it 'detects corrupt chunks' do
    result = described_class.validate('{([(<{}[<>[]}>{[]{[(<()>')
    expect(result[:result]).to eq(:corrupted)
    expect(result[:char]).to eq('}')
  end

  it 'detects valid chunks' do
    expect(described_class.validate('(())')[:result]).to eq(:valid)
  end
end
