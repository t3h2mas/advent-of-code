require_relative './lib/bits'

binary_strings = File.read('./input.txt').split("\n").filter { |l| !l.nil? && l.length > 1 }

gamma = Bits.gamma(binary_strings)

epsilon = Bits.epsilon(gamma)

puts "Part one: #{gamma * epsilon}"

oxygen_rating = Bits.oxygen(binary_strings)
scrubber_rating = Bits.scrubber(binary_strings)

puts "Part two: #{oxygen_rating * scrubber_rating}"
