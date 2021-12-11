class Vents
  def initialize(segments)
    @points = {}

    segments.each do |segment|
      start_x, start_y = segment.first
      end_x, end_y = segment.last

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
