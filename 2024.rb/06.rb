input = <<~TEXT
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
TEXT

input = File.read('./06_input.txt')

grid = Grid.new(input.split("\n"))

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

  def value_at(cell)
    @grid[cell.y][cell.x] 
  end
end

player = grid.locations_of('^').first
direction = :N

turns = { N: :E, E: :S, S: :W, W: :N }

locs = Set[{ pos: player, dir: direction }]

walls = Set.new

# viz = input.dup.split("\n")

while player.on?(grid)
  next_spot = player.move(direction)

  if !next_spot.on?(grid)
    break
  elsif grid.value_at(next_spot) == '#'
    walls << next_spot
    direction = turns[direction]
  else
    player = next_spot
    locs << { pos: player, dir: direction };
  end
  # locs.each { |cell| viz[cell.y][cell.x] = '!' }
  # puts viz.join("\n")
  # puts "\n"
end

locs

locs.map { _1[:pos] }.uniq.count

locs.count

walls

locs.count

obstructions = Set.new

potential_obstructions = Set.new

run_ins = Set.new

locs.each do |loc|
  next_spot = loc[:pos].move(loc[:dir])
  next if locs.include?(next_spot)
  next unless next_spot.on?(grid)
  next if grid.value_at(next_spot) == '#'

  # pretend to turn and walk until there is obstruction or wall
  turn_dir = turns[loc[:dir]]
  walker = loc[:pos].dup

  puts "current loc: #{loc}, trying direction: #{turn_dir}"

  while walker.on?(grid)
    walker_next = walker.move(turn_dir)
    puts "walker: #{walker}"
    break unless walker_next.on?(grid)
    if grid.value_at(walker_next) == '#' #&& locs.include?({ pos: walker, dir: turns[turn_dir] })
      run_ins << { pos: walker, dir: turn_dir }
      potential_obstructions << next_spot
      break
    end

    walker = walker_next
  end
end;

run_ins.map { _1[:pos] }.uniq.count

run_ins.each do |ri|
  loc = locs.filter { _1[:pos] == ri[:pos] }
  next if loc.empty?
  p ri
  p locs
end;
