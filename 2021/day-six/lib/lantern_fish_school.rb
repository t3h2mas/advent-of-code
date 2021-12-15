class LanternFishSchool

  def initialize(founding_fish)
    @fish = Hash.new(0)
    founding_fish.each { |life_cycle| @fish[life_cycle] += 1 }
  end

  def cycle(days)
    (0...days).each { cycle_day }
  end

  def count
    @fish.values.sum
  end

  private

  def cycle_day
    parents = @fish[0]
    (1..8).each do |timer|
      @fish[timer - 1] = @fish[timer]
    end
    @fish[8] = parents
    @fish[6] += parents
  end
end
