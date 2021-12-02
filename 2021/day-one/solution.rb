require_relative './lib/sonar'
require_relative './lib/advent_io'

measurements = AdventIO.as_lines(File.new('./input.txt', 'r')).map(&:to_i)

puts 'Part one:'
puts Sonar.depth_increases(measurements)

puts 'Part two:'
puts Sonar.depth_increases_over_window(measurements)
