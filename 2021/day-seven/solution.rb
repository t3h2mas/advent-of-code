# frozen_string_literal: true

horizontal_positions = [16, 1, 2, 0, 4, 2, 7, 1, 2, 14]

def solve(positions)
  xs = positions.sort
  cost = {}

  for i in (0...xs.length)
    cost[i] = 0
    for j in (0...xs.length)
      next if i == j

      if j  > i
        cost[i] += xs[j] - xs[i]
      else
        cost[i] += xs[i] - xs[j]
      end
    end
  end

  cost.values.min
end

def solve2(xs)
  costs = []
  (xs.min..xs.max).each do |e|
    cost = 0
    xs.each do |x|
      cost += travel_cost(x, e)
    end
    costs << cost
  end
  costs.min
end

def travel_cost(x, y)
  n = (x - y).abs
  (n * (n + 1)) / 2
end

puts 'part one'
puts solve(horizontal_positions)
puts 'Not there yet' unless solve(horizontal_positions) == 37

input = File.read('./input.txt').split(',').filter { |l| !l.empty? }.map(&:to_i)

puts solve(input)

cost = travel_cost(1, 5)
raise "expected 10, got #{cost}" unless cost == 10

cost = travel_cost(0, 5)
raise "expected 15, got #{cost}" unless cost == 15

puts 'part two'
puts solve2(horizontal_positions)
puts 'Not there yet' unless solve2(horizontal_positions) == 168
puts solve2(input)
