require 'lantern_fish'

RSpec.describe LanternFish do
  describe '#school' do
    it 'lists 26 fish after 18 days' do
      lantern_fish = LanternFish.new([3, 4, 3, 1, 2])

      lantern_fish.cycle(18)

      expect(lantern_fish.school.count).to eq(26)
    end
  end
end
