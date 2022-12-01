require 'surveyor'

RSpec.describe Surveyor do
  it 'finds low points' do
    cases = [
      {
        input: [[1, 1, 1], [1, 1, 1], [1, 1, 1]],
        expected: [],
        name: 'no low points'
      },
      {
        input: [[4, 4, 4], [5, 2, 6], [3, 3, 3]],
        expected: [2],
        name: 'one low point'
      },
      {
        input: [[5, 9, 9, 5],
                [9, 9, 9, 9],
                [9, 9, 9, 9],
                [5, 9, 9, 5]],
        expected: [5, 5, 5, 5],
        name: 'redditor support'
      }
    ]

    cases.each do |c|
      expected = c[:expected]
      name = c[:name]

      s = Surveyor.new(c[:input])

      got = s.heights(s.low_points)

      expect(got).to eq(expected), "#{name}: expected #{expected}, got #{got}"
    end
  end

  it 'finds the size of a basin given a point' do
    cases = [
      {
        grid: [[9, 9, 9], [9, 2, 9], [9, 9, 9]],
        input: Point.new(row: 1, col: 1),
        expected: 1,
        name: 'basin with the size of 1'
      },
      {
        grid: [[9, 9, 9], [9, 2, 9], [9, 1, 9]],
        input: Point.new(row: 1, col: 1),
        expected: 2,
        name: 'basin with the size of 2'
      }
    ]

    cases.each do |c|
      expected = c[:expected]
      name = c[:name]

      got = Surveyor.new(c[:grid]).basin_size(c[:input])

      expect(got).to eq(expected), "#{name}: expected #{expected}, got #{got}"

    end
  end
end
