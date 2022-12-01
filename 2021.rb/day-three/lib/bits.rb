module Bits
  def self.gamma(binary_strings)
    counters = []
    binary_strings[0].each_char { counters << [0, 0] }

    binary_strings[0].split('').each_index do |i|
      binary_strings.each do |s|
        counters[i][0] += 1 if s[i] == '0'
        counters[i][1] += 1 if s[i] == '1'
      end
    end

    counters.map do |counter|
      if counter[0] > counter[1]
        '0'
      else
        '1'
      end
    end.join('').to_i(2)
  end

  def self.epsilon(gamma)
    ones = gamma.to_s(2).split('').map { |_| '1' }.join('').to_i(2)
    ones ^ gamma
  end

  def self.oxygen(binary_strings)
    find_sensor(binary_strings, :>, '1')
  end

  def self.scrubber(binary_strings)
    find_sensor(binary_strings, :<, '0')
  end

  def self.find_sensor(binary_strings, method, fallback)
    results = binary_strings.clone

    (0..results[0].length).each do |i|
      break if results.length == 1

      new_results = []

      target = target_bit(results, method, i, fallback)

      results.each do |s|
        new_results << s if s[i] == target
      end
      results = new_results
    end

    results[0].to_i(2)
  end

  def self.target_bit(results, method, index, fallback)
    occurrence = count_row_occurences(results, index)

    if occurrence['1'].send(method, occurrence['0'])
      '1'
    elsif occurrence['0'].send(method, occurrence['1'])
      '0'
    else
      fallback
    end
  end

  def self.count_row_occurences(binary_strings, index)
    counter = { '0' => 0, '1' => 0 }

    binary_strings.each do |s|
      counter[s[index]] += 1
    end

    counter
  end
end
