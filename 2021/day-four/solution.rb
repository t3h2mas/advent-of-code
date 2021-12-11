require_relative 'lib/bingo'

input_chunks = File.read('./input.txt').split("\n\n")

to_call = input_chunks[0].split(',').map(&:to_i)

cards = input_chunks.slice(1..).map do |chunk|
  chunk.split("\n").map { |line| line.split(' ').map(&:to_i) }
end

puts 'Day four part one'
game = Bingo.new(cards)

to_call.each do |number|
  game.mark(number)

  if game.won?
    puts "solution: #{game.winner.unmarked.sum * number}"
    break
  end
end
