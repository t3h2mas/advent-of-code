require_relative 'point'

require 'set'
require 'pry'

# Surveys caves
class Surveyor
  def initialize(cave)
    @cave = cave
  end

  def low_points
    low_points = []

    @cave.each_index do |row|
      @cave[row].each_index do |col|
        point = Point.new(row: row, col: col)

        low_points << point if low?(point)
      end
    end

    low_points
  end

  def heights(points)
    points.map(&method(:height_of))
  end

  def low?(point)
    valid_neighbors(point).all? do |neighbor|
      neighbor_height = height_of(neighbor)

      height_of(point) < neighbor_height
    end
  end

  def height_of(point)
    @cave.dig(point.row, point.col)
  end

  def valid_neighbors(point)
    point.neighbors.filter(&method(:point_in_range?))
  end

  def point_in_range?(point)
    point.row.between?(0, @cave.length - 1) && point.col.between?(0, @cave[0].length - 1)
  end

  def basin_size(point)
    visited = Set.new
    to_visit = []

    to_visit << point

    size = 0

    until to_visit.empty?
      p = to_visit.pop

      next if visited.include?(p)

      visited << p

      next if height_of(p) == 9

      size += 1

      valid_neighbors(p).each { |n| to_visit << n }
    end

    size
  end
end
