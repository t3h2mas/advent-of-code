defmodule Five do
  @moduledoc """
  Advent of Code, day five
  """

  @input_file "five.txt"

  def read_input_file do
    case File.read(@input_file) do
      {:ok, body} ->
        body
      {:error, reason} ->
        raise reason
    end
  end

  @doc """
  Parse columnar list string `jumps` into list of integers
  """
  def parse_jumps(jumps) do
    jumps
    |> String.split("\n", trim: :true)
    |> Enum.map(fn x ->
      x
      |> String.trim
      |> String.to_integer
    end)
  end

  defp cond_inc_or_dec(val) do
    case val > 2 do
      :true -> val - 1
      :false -> val + 1
    end
  end

  defp always_inc(x) do
    x + 1
  end

  def cycle_jumps(jumps, inc_f) do
    start_index = 0
    step = 0

    cycle_jumps(jumps, start_index, step, inc_f)
  end

  def cycle_jumps(jumps, index, step, inc_f) do
    case Enum.fetch(jumps, index) do
      {:ok, pointer_val} ->
        next_jumps = List.replace_at(jumps, index, inc_f.(pointer_val))
        cycle_jumps(next_jumps, index + pointer_val, step + 1, inc_f)
      :error -> step
    end
  end

  def solve_part_one do
    read_input_file() |> solve_part_one
  end

  @doc """
  solve part one by passing `always_inc`/1 as the increment func
  """
  def solve_part_one(str) do 
    str
    |> parse_jumps
    |> (fn x ->
          cycle_jumps(x, &always_inc/1)
        end).()
  end

  def solve_part_two do
    read_input_file() |> solve_part_two
  end

  @doc """
  solve part two by passing `cond_inc_or_dec`/1 as the increment func
  """
  def solve_part_two(str) do 
    str
    |> parse_jumps
    |> (fn x ->
          cycle_jumps(x, &cond_inc_or_dec/1)
    end).()
  end
end
