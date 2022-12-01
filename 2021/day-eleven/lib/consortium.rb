require 'pry'
require 'set'

class Consortium
  attr_reader :flashes

  def initialize(input)
    @flashes = 0

    lines = input.split("\n").filter { |l| !l.empty? }
    @consortium = lines.map.with_index do |row, y|
      row.split('').map.with_index do |col, x|
        Octopus.new(x, y, col.to_i, &method(:handle_flash))
      end
    end

    @to_flash = []
    @flashed = Set.new
    @last_flashed = 0
  end

  def cycle(cycles)
    cycles.times { run_cycle }
    display
    self
  end

  def detect_synchronous_flash
    required_flashes = @consortium.size * @consortium[0].size
    step = 0

    until @last_flashed == required_flashes
      run_cycle
      step += 1
    end

    step
  end

  def display
    @consortium.each do |row|
      row.each do |col|
        print "#{col.value} "
      end
      puts
    end
    puts
    nil
  end

  private

  def run_cycle
    increment_all
    handle_flashes
    @flashed.clear
  end

  def handle_flashes
    until @to_flash.empty? do
      octopus = @to_flash.pop

      next if @flashed.include?(octopus)

      @flashed << octopus

      neighbors(octopus).each(&:increment)

      @flashes += 1
    end

    @flashed.each(&:reset)

    @last_flashed = @flashed.count
  end

  def increment_all
    @consortium.each do |row|
      row.each(&:increment)
    end
  end

  def handle_flash(octopus)
    @to_flash << octopus
  end

  def neighbors(octopus)
    res = []
    (-1..1).each do |x|
      (-1..1).each do |y|
        next if x.zero? && y.zero?

        neighbor_x = octopus.x + x
        neighbor_y = octopus.y + y

        next if neighbor_x.negative? || neighbor_y.negative?

        next if neighbor_x >= @consortium[0].size || neighbor_y >= @consortium.size

        res << @consortium[neighbor_y][neighbor_x]
      end
    end

    res
  end

  class Octopus
    attr_reader :x, :y, :value

    def initialize(x, y, value, &notify_flashable)
      @x = x
      @y = y
      @value = value
      @notify_flashable = notify_flashable
    end

    def increment
      @value += 1

      return unless @value > 9

      @notify_flashable.call(self)
    end

    def reset
      @value = 0
    end
  end
end
