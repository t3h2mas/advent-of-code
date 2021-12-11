require_relative './lib/vents'
require_relative './lib/segments'

input = File.read('./input.txt')

segments = Segments.parse(input, Segments.method(:validate_without_diagonals))
vents = Vents.new(segments)

puts "Day five part one: #{vents.overlapping}"
