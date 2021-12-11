require_relative 'lib/bingo'

input_chunks = File.read('./input.txt').split("\n\n")

to_call = input_chunks[0].split(',').map(&:to_i)

cards = input_chunks.slice(1..).map do |chunk|
  chunk.split("\n").map { |line| line.split(' ').map(&:to_i) }
end

puts 'Day four part one'
game = Bingo.new(cards)

to_call.each { |number| game.mark(number) }

first_win = game.wins.first
last_win = game.wins.last

puts "solution: #{first_win[:card].unmarked.sum * first_win[:winning_number]}"

puts 'Day four part two'

puts "solution: #{last_win[:card].unmarked.sum * last_win[:winning_number]}"
