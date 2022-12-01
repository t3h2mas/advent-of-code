require_relative 'lib/surveyor'

def as_grid(input_string)
  input_string.split("\n").map { |row| row.split('').map(&:to_i) }
end

def try_solution(fname)
  puts "trying for file: #{fname}"
  grid = as_grid(File.read(fname))

  s = Surveyor.new(grid)
  low_points = s.low_points
  risk_levels = s.heights(low_points).map { |p| p + 1 }
  puts "total risk: #{risk_levels.sum}"

  largest_basin_sizes = low_points
    .map { |lp| s.basin_size(lp) }
    .sort
    .reverse
    .slice(0...3)

  puts "largest basins: #{largest_basin_sizes}"
  puts largest_basin_sizes.reduce(&:*)
end

try_solution('./small_input.txt')
try_solution('./input.txt')

