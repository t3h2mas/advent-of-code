class Vents
  def initialize(segments)
    @points = {}

    segments.each do |segment|
      x1, y1 = segment.first
      x2, y2 = segment.last

      x = x1
      y = y1

      loop do
        key = { x: x, y: y }
        @points[key] = @points.fetch(key, 0) + 1

        break if x == x2 && y == y2

        if x1 < x2
          x += 1
        elsif x1 > x2
          x -= 1
        end

        if y1 < y2
          y += 1
        elsif y1 > y2
          y -= 1
        end
      end
    end
  end

    def overlapping
      @points.values.filter { |n| n > 1 }.count
    end
  end
