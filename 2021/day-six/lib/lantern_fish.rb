class LanternFish

  attr_reader :school

  def initialize(school)
    @school = school
  end

  def cycle(days)
    (0...days).each {  cycle_day }
  end

  def cycle_day
    new_school = @school.clone
    @school.each_with_index do |fish_timer, index|
      if fish_timer.zero?
        new_school << 8
        new_school[index] = 6
      else
        new_school[index] -= 1
      end

      @school = new_school
    end
  end
end
