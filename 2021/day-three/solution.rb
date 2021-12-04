require_relative './lib/bits'

binary_strings = File.read('./input.txt').split("\n").filter { |l| !l.nil? && l.length > 1 }

gamma = Bits.gamma(binary_strings)

epsilon = Bits.epsilon(gamma)

puts "Part one: #{gamma * epsilon}"
