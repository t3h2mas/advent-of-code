input = 'xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))'

input = File.read('03_input.txt')

input.scan(/mul\(([0-9]{1,3},[0-9]{1,3})\)/).flatten.map { _1.split(',').map(&:to_i) }.reduce(0) { |acc, el| acc += el[0] * el[1] }

# part two
# input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
#
# input = "do()xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
#
input = "do()xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))don't()asdfmul(3,2)"

res = 0
start = 0
stop = input.index("don't()")

slice = input[start..stop]
puts "starting slice: #{slice}"

until slice.empty?
  puts "slice: #{slice}, start: #{start}, stop: #{stop}"
  slice.scan(/mul\(([0-9]{1,3},[0-9]{1,3})\)/).flatten.map { _1.split(',').map(&:to_i) }.each { p _1; res += _1[0] * _1[1] }

  break if stop.nil?

  start = input.index("do()", stop)
  stop = input.index("don't()", start)

  slice = input[start..stop]
end

p res
