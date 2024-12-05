word_search = <<~INPUT
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
INPUT

word_search = File.read('./04_input.txt');

COMPASS = {
  NW: { x: -1, y: -1 },
  N: { x: 0, y: -1 },
  NE: { x: 1, y: -1 },
  W: { x: -1, y: 0 },
  E: { x: 1, y: 0 },
  SW: { x: -1, y: 1 },
  S: { x: 0, y: 1 },
  SE: { x: 1, y: 1 }
}

def move(loc, direction)
  raise ArgumentError, "invalid direction: #{direction}" unless COMPASS.include?(direction)

  {
    x: loc[:x] + COMPASS[direction][:x],
    y: loc[:y] + COMPASS[direction][:y]
  }
end

def in_bounds?(grid, loc)
  loc[:x].between?(0, grid.first.length - 1) && loc[:y].between?(0, grid.length - 1)
end

def completes_word?(grid, loc, letters, direction)
  return true if letters.empty?
  return false unless in_bounds?(grid, loc)

  return false unless grid[loc[:y]][loc[:x]] == letters[0]

  completes_word?(grid, move(loc, direction), letters[1..], direction)
end

def begins_words(grid, loc, word)
  return 0 unless grid[loc[:y]][loc[:x]] == word[0]

  COMPASS.keys.filter { |direction| completes_word?(grid, move(loc, direction), word.chars[1..], direction) }.count
end

grid = word_search.split("\n")

def find_char(grid, needle)
  locs = []

  grid.each_with_index do |line, y|
    line.chars.each_with_index do |char, x|
      locs << { x:, y: } if char == needle
    end
  end

  locs
end

locs = find_char(grid, 'X')

locs.sum { begins_words(grid, _1, "XMAS") }

# part two
grid = word_search.split("\n")

locs = find_char(grid, 'A')

def begins_x_words(grid, loc, word)
  return false unless grid[loc[:y]][loc[:x]] == word[1] # assume 3 chars

  corner_lines = [[:NW, :SE], [:SW, :NE]].map { |line| line.map { move(loc, _1) } }

  return false unless corner_lines.flatten.all? { in_bounds?(grid, _1) }

  values = corner_lines.map { |line| line.map { grid[_1[:y]][_1[:x]] }.to_set }

  return values.all? { _1 == Set[word[0], word[2]] } # keep assuming 3 chars
end

locs.filter { begins_x_words(grid, _1, "MAS") }.count
