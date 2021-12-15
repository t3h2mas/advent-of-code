require_relative 'lib/lantern_fish_school'

input = File.read('./input.txt').split(',').map(&:to_i)

puts 'Day six part one: '

fish = LanternFishSchool.new(input)

fish.cycle(80)

puts fish.count

puts 'Day six part two: '

fish.cycle(256 - 80)

puts fish.count
