def slope(p1, p2)
  x1, y1 = p1
  x2, y2 = p2

  ((y2.to_f - y1.to_f) / (x2.to_f - x1.to_f)).abs
rescue ZeroDivisionError
  0
end

class Segments
  def self.parse(input, validator)
    point_pairs = input.split("\n").map { |line| line.split(' -> ').map { |pair| pair.split(',').map(&:to_i) } }

    point_pairs.filter { |segment| validator.call(segment) }
  end

  def self.validate_without_diagonals(segment)
    start_x, start_y = segment.first
    end_x, end_y = segment.last

    !(start_x != end_x && start_y != end_y)
  end

  def self.validate_with_diagonals(segment)
    #x1, y1 = segment.first
    #x2, y2 = segment.last

    #(x1 != x2 && y1 != y2) ||
    #(Math.atan2((y2 - y1).abs, (x2 - x1).abs) * 180 / Math::PI) == 45

    #!(x1 != x2 && y1 != y2) ||
      #slope(segment.first, segment.last) == 1
    s = slope(segment.first, segment.last)

    validate_without_diagonals(segment) || diagonal?(segment)
  end

  def self.validate_all(segment)
    true
  end

  def self.diagonal?(segment)
    slope(segment.first, segment.last) == 1
  end
end
