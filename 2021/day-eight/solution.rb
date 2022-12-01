require 'set'

text = File.read('./input.txt')

segments = text.split("\n").map do |l|
  parts = l.split(' | ')
  { input: parts[0], output: parts[1] }
end

segments_for_digit = {
  2 => 1,
  4 => 4,
  3 => 7,
  7 => 8
}

counts = Hash.new(0)

segments.map { |s| s[:output] }.each do |output|
  output.split(' ').each do |digit|
    digit = segments_for_digit[digit.length]
    counts[digit] += 1 unless digit.nil?
  end
end

puts "Day one part one"
puts counts.values.sum

entry = {:input=>"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb",
           :output=>"fdgacbe cefdb cefbgd gcbe"}

def digits_from(entry)
  entry_input = entry[:input]

  signals = entry_input.split(' ').map { |s| Set.new(s.split('')) }
  signal_groups = signals.group_by(&:length)


  digits = {}

  # digit: length; 1:2, 4:4, 7:3, 8:7
  # use variables for future filtering
  one = signal_groups[2].first
  four = signal_groups[4].first
  seven = signal_groups[3].first
  eight = signal_groups[7].first
  digits[one] = 1
  digits[four] = 4
  digits[seven] = 7
  digits[eight] = 8


  # deduce nine
  # has everything in common with 1,4,7 - has 6 signals
  nine_includes = one.clone.merge(four).merge(seven)
  nine_signals = signal_groups[6].find { |group| nine_includes <= group }
  digits[nine_signals] = 9

  signal_groups[6].delete(nine_signals)

  # signal_groups[6] still contains 0 and 6

  zero_signals = signal_groups[6].find { |group| one <= group }
  digits[zero_signals] = 0

  signal_groups[6].delete(zero_signals)

  digits[signal_groups[6].first] = 6

  # signal_groups[5] contains 2, 3, 5

  three_signals = signal_groups[5].find { |group| one <= group }
  digits[three_signals] = 3

  signal_groups[5].delete(three_signals)

  # signal_groups[5] contains 2, 5

  # 5 has a difference of 1 from 9, 2 has a difference of 2 from 9
  digits[signal_groups[5].find { |group| (nine_signals - group).length == 1 }] = 5

  digits[signal_groups[5].find { |group| (nine_signals - group).length == 2 }] = 2

  digits
end

total = 0
segments.each do |s|
  digits = digits_from(s)

  total += s[:output].split(' ').map { |g| Set.new(g.split('')) }.map { |signals| digits[signals] }.join('').to_i
end

puts "total: #{total}"
