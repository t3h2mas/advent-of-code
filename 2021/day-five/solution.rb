require_relative './lib/vents'

input = File.read('./input.txt')

vents = Vents.new(input)

puts "Day five part one: #{vents.overlapping}"
