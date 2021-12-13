require_relative './lib/vents'
require_relative './lib/segments'

input = File.read('./input.txt')

count = input.split("\n").count

puts "lines: #{count}"

def part_one(input)
  segments = Segments.parse(input, Segments.method(:validate_without_diagonals))
  vents = Vents.new(segments)

  puts "Day five part one: #{vents.overlapping}"
  puts "Segments: #{segments.count}"
end

def part_two(input)
  segments = Segments.parse(input, Segments.method(:validate_with_diagonals))
  vents = Vents.new(segments)

  puts "Day five part two: #{vents.overlapping}"
  puts "Segments: #{segments.count}"
end

part_one(input)
part_two(input)
