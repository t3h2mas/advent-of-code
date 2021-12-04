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

  def self.oxygen(numbers)
    
  end
end
