require_relative 'lib/consortium'

puts 'Flash count after 100:'
puts Consortium.new(File.read('./input.txt')).cycle(100).flashes

puts 'Synchronized at step:'
puts Consortium.new(File.read('./input.txt')).detect_synchronous_flash
