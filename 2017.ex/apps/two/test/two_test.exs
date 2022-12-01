defmodule TwoTest do
  use ExUnit.Case

  test "mapify takes string, returns 2d array" do
    assert Two.mapify("1\t2\n3\t4\n") == [[1, 2], [3, 4]]
  end

  test "checksum(cells) returns the 'checksum' where length(cells) is 1" do
    seed = [[5, 4, 3]]
    assert Two.checksum(seed) == 2
  end
  test "checksum(cells) returns the 'checksum' where length(cells) is > 1" do
    seed = [[5, 4, 3], [10, 1, 2]]
    assert Two.checksum(seed) == 11
  end

  test "permutations/1 makes a list of possible combinations as tuples" do
    assert Two.permutations([1, 2]) == [{1, 1}, {1, 2}, {2, 1}, {2, 2}]
  end
end
