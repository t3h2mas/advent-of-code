class Segments
  def self.parse(input, validator)
    point_pairs = input.split("\n").map { |line| line.split(' -> ').map { |pair| pair.split(',').map(&:to_i) }.sort }

    point_pairs.filter { |segment| validator.call(segment) }
  end

  def self.validate_without_diagonals(segment)
    start_x, start_y = segment.first
    end_x, end_y = segment.last

    !(start_x != end_x && start_y != end_y)
  end
end
