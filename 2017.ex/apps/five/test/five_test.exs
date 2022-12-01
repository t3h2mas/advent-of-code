defmodule FiveTest do
  use ExUnit.Case

  test "jumps_to_list returns list from string" do
    assert Five.parse_jumps("0 \n3 \n0 \n1 \n-3") == [0, 3, 0, 1, -3]
  end

  test "solve_part_one" do
    assert Five.solve_part_one("0 \n3 \n0 \n1 \n-3") == 5
  end
  test "solve_part_two" do
    assert Five.solve_part_two("0 \n3 \n0 \n1 \n-3") == 10
  end
end
