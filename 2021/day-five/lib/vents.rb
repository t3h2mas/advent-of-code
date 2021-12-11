class Vents
  def initialize(input)
    @points = {}

    point_pairs = input.split("\n").map { |line| line.split(' -> ').map { |pair| pair.split(',').map(&:to_i) }.sort }

    point_pairs.each do |pair|
      start_point = pair.first
      end_point = pair.last

      start_x = start_point[0]
      start_y = start_point[1]

      end_x = end_point[0]
      end_y = end_point[1]

      next if start_x != end_x && start_y != end_y

      (start_x..end_x).each do |segment_x|
        (start_y..end_y).each do |segment_y|
          key = { x: segment_x, y: segment_y }

          @points[key] = @points.fetch(key, 0) + 1
        end
      end
    end
  end

  def overlapping
    @points.values.filter { |n| n > 1 }.count
  end
end
