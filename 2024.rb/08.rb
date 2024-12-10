input = <<~TEXT
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
TEXT

input = <<~TEXT
..........
..........
..........
....a.....
..........
.....a....
..........
..........
..........
..........
TEXT

input = File.read('./08_input.txt')

class Cell
  attr_reader :x, :y

  def initialize(x:, y:)
    @x = x
    @y = y
  end

  def +(other)
    self.class.new(x: self.x + other.x, y: self.y + other.y)
  end

  def move(direction)
    raise ArgumentError, "invalid direction: #{direction}" unless compass.include?(direction)

    self.+(compass[direction])
  end

  def on?(grid)
    grid.include?(self)
  end

  def eql?(other)
    self.x == other.x && self.y == other.y
  end

  def hash
    [self.x, self.y].hash 
  end

  def compass
    @compass ||= {
      NW: Cell.new(x: -1, y: -1),
      N: Cell.new(x: 0, y: -1),
      NE: Cell.new(x: 1, y: -1),
      W: Cell.new(x: -1, y: 0),
      E: Cell.new(x: 1, y: 0),
      SW: Cell.new(x: -1, y: 1),
      S: Cell.new(x: 0, y: 1),
      SE: Cell.new(x: 1, y: 1)
    }
  end

  def inspect
    "<Cell x: #{x}, y: #{y}>"
  end

  def to_s
    "<Cell x: #{x}, y: #{y}>"
  end
end

class Grid
  def initialize(grid)
    @grid = grid
  end

  def include?(cell)
    cell.x.between?(0, @grid.first.length - 1) && cell.y.between?(0, @grid.length - 1)
  end

  def locations_of(needle)
    locs = []

    @grid.each_with_index do |line, y|
      line.chars.each_with_index do |char, x|
        locs << Cell.new(x:, y:) if char == needle
      end
    end

    locs
  end

  def each_cell(&block)
    @grid.each_with_index do |line, y|
      line.chars.each_with_index do |_, x|
        block.call(Cell.new(x:, y:))
      end
    end
  end

  def value_at(cell)
    @grid[cell.y][cell.x] 
  end
end

grid = Grid.new(input.split("\n"))

antennas = {}

grid.each_cell do |cell|
  val = grid.value_at(cell)
  next if val == '.'
  
  antennas[val] = Set.new unless antennas.include?(val)

  antennas[val] << cell
end

antennas

antinodes = Set.new

seen = Set.new

antennas.each do |(frequency, locations)|
  # puts "freq: #{frequency}, locations: #{locations}"
  locations.each do |location|
    locations.each do |other|
      next if location.eql?(other)
      next if seen.include?(Set[location, other])
      seen << Set[location, other]

      rise = location.y - other.y
      run = location.x - other.x

      # puts "start: #{location}, other: #{other}, rise: #{rise}, run: #{run}"

      fst = location + Cell.new(x: run, y: rise)
      snd = other + Cell.new(x: -run, y: -rise)

      antinodes << fst if fst.on?(grid)
      antinodes << snd if snd.on?(grid)
    end
  end
  # puts "\n"
end;

antinodes

antinodes.length

viz = input.dup.split("\n")

antinodes.each do |c|
  viz[c.y][c.x] = '#'
end

puts viz.join("\n")

# part two

antennas.each do |(frequency, locations)|
  locations.each do |location|
    antinodes << location
    locations.each do |other|
      next if location.eql?(other)
      next if seen.include?(Set[location, other])
      seen << Set[location, other]

      rise = location.y - other.y
      run = location.x - other.x

      fst = location + Cell.new(x: run, y: rise)

      while fst.on?(grid)
        antinodes << fst
        fst = fst + Cell.new(x: run, y: rise)
      end

      snd = other + Cell.new(x: -run, y: -rise)

      while snd.on?(grid)
        antinodes << snd if snd.on?(grid)
        snd = snd + Cell.new(x: -run, y: -rise)
      end
    end
  end
end;

antinodes.count
