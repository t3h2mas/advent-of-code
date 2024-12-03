input = <<~INPUT
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
INPUT

input = File.read('./02_input.txt')

def level_safe?(level)
  op = if level[0] == level[1]
                nil
              elsif level[0] > level[1]
                :>
              else
                :<
              end

  return false if op.nil?

  (level.length - 1).times do |idx|
    return false unless level[idx].send(op, level[idx + 1]) 
    return false unless (1..3).cover?((level[idx] - level[idx + 1]).abs)
  end

  true
end

# level_safe2? wasn't pretty.
def level_safe3?(level)
  cur_level = level

  # could be faster if we started at the first offending position
  delete_at = 0

  while delete_at <= level.length do
    return true if level_safe?(cur_level)

    cur_level = level.filter.with_index { _2 != delete_at }
    delete_at += 1
  end

  false
end

p input.split("\n").map { _1.split(" ").map(&:to_i) }.filter { level_safe3?(_1) }.length;
