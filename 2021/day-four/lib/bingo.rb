require 'set'

class Bingo

  attr_reader :winner

  def initialize(cards)
    @cards = cards.map { |c| Card.new(c) }
    @winner = nil
  end

  def mark(number)
    @cards.each do |card|
      card.mark(number)
      if card.won?
        @winner = card
        break
      end
    end
  end

  def won?
    @winner != nil
  end
end

class Card
  def initialize(numbers)
    @numbers = numbers

    @marked = Set.new
    @lookup = {}
    @won = false

    @numbers.each_with_index do |row, row_idx|
      row.each_with_index do |val, col_idx|
        @lookup[val] = { row: row_idx, col: col_idx }
      end
    end
  end

  def won?
    @won
  end

  def mark(number)
    position = @lookup[number]

    return if position.nil?

    @marked << number

    check_for_wins(position)
  end

  def check_for_wins(position)
    won = true
    # check for horizontal win
    (0...@numbers[0].length).each do |col|
      next if @marked.include?(@numbers[position[:row]][col])

      won = false
      break
    end

    if won
      @won = true
      return
    end

    won = true
    # check for vertical win
    (0...@numbers.length).each do |row|
      next if @marked.include?(@numbers[row][position[:col]])

      won = false
      break
    end

    @won = true if won
  end

  def unmarked
    result = []

    @numbers.each do |row|
      row.each do |number|
        result << number unless @marked.include?(number)
      end
    end

    result
  end
end

