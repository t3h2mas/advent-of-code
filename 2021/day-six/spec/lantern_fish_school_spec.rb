require 'lantern_fish_school'

RSpec.describe LanternFishSchool do
  describe 'life' do
    it 'lists 26 fish after 18 days' do
      lantern_fish = described_class.new([3, 4, 3, 1, 2])

      lantern_fish.cycle(18)

      expect(lantern_fish.count).to eq(26)
    end
  end
end
