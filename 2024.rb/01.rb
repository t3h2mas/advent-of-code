location_ids = <<~INPUT
3   4
4   3
2   5
1   3
3   9
3   3
INPUT

location_ids = File.read('./01_input.txt')

puts location_ids

lists = location_ids.split("\n").flat_map { _1.split('  ').map(&:to_i) }.partition.with_index { |_, idx| idx.even? }.map(&:sort)


delta = 0

lists.first.length.times { |i| delta += (lists[0][i] - lists[1][i]).abs } 

delta

# part two
left = lists[0]
right = lists[1]
counts = right.tally

left.reduce(0) { |acc, n| acc += n * counts.fetch(n, 0) }
