require_relative 'lib/lantern_fish'

input = File.read('./input.txt').split(',').map(&:to_i)

puts 'Day six part one: '

fish = LanternFish.new(input)

fish.cycle(80)

puts fish.school.count

puts 'Day six part two: '

fish.cycle(256 - 80)

puts fish.school.count
