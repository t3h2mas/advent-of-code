# Represents a 2d point
class Point

  attr_reader :row, :col

  def initialize(row:, col:)
    @row = row
    @col = col
  end

  def neighbors
    [
      modify(row: row + 1), # above
      modify(col: col - 1), # left
      modify(row: row - 1), # below
      modify(col: col + 1)  # right
    ]
  end

  def modify(**args)
    Point.new(**{ row: row, col: col }.merge(args))
  end

  def to_s
    "<Point: row: #{row}, col: #{col}>"
  end

  def eql?(other)
    row == other.row && col == other.col
  end

  def hash
    "#{row}-#{col}".hash
  end
end
