class Navigation
  attr_reader :horizontal_position, :depth

  def initialize
    @horizontal_position = 0
    @depth = 0
  end
end

class DirectionalNavigation < Navigation
  def move_forward(count)
    @horizontal_position += count
  end

  def move_backward(count)
    @horizontal_position -= count
  end

  def move_up(count)
    @depth -= count
  end

  def move_down(count)
    @depth += count
  end
end

class AimNavigation < Navigation
  def initialize
    super
    @aim = 0
  end

  def move_forward(count)
    @horizontal_position += count
    @depth += @aim * count
  end

  def move_backward(count)
    @horizontal_position -= count
  end

  def move_up(count)
    @aim -= count
  end

  def move_down(count)
    @aim += count
  end
end

class Submarine
  def initialize(navigation)
    @navigation = navigation
  end

  def move(direction, count)
    case direction
    when :forward
      @navigation.move_forward(count)
    when :backward
      @navigation.move_backward(count)
    when :up
      @navigation.move_up(count)
    when :down
      @navigation.move_down(count)
    end
  end

  def horizontal_position
    @navigation.horizontal_position
  end

  def depth
    @navigation.depth
  end
end
