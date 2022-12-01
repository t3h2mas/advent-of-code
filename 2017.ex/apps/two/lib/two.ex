defmodule Two do
  @moduledoc """
  Advent of code, day 2
  """

  @input_file "input.txt"

  @doc """
  Take the spreadsheet string, give the nested array form

  ## Parameters
    - s: String that represents the speadsheet input
  """
  def mapify(s) do
    s
    |> String.split("\n", trim: :true)
    |> Enum.map(fn x -> 
         Enum.map(String.split(x, "\t"), fn x -> String.to_integer(x) end) 
       end)
  end

  @doc """
  Checksum/0 Parses input file, passes body to checksum/1
  """
  def checksum do
    # read in the file
    case File.read(@input_file) do
      {:ok, body} ->
        checksum(mapify(body))
      {:error, reason} ->
        {:error, reason}
    end
  end

  @doc """
  checksum/1 
   Takes the spreadsheet as an array
   Gives the 'checksum' as an integer
  """
  def checksum(cells) do
    cells 
    |> Enum.map(fn arr -> 
        max = Enum.max(arr)
        min = Enum.min(arr)
        max - min
      end)
    |> Enum.reduce(fn(cur, acc) -> cur + acc end)
  end

  # part two
  
  def permutations(arr) do
    arr 
    |> Enum.flat_map(fn x -> 
      Enum.map(arr, fn y ->
        max = Enum.max([x, y])
        min = Enum.min([x, y])
        {max, min}
      end)
    end)
    |> Enum.filter(fn {x, y} -> x != y end)
    |> Enum.uniq
  end

  def sum_divisible(arr) do
    arr
    |> Enum.reduce(0, fn({x, y}, acc) ->
      case rem(x, y) == 0 do
        :true -> acc + (x / y)
        :false -> acc
      end
    end)
  end

  def checksum_line(line) do
    line
    |> permutations
    |> sum_divisible
    |> round
  end
  
  def checksum2(cells) do
    cells 
    |> Enum.map(&checksum_line/1)
    |> Enum.sum
  end

  def solve_two do
    case File.read(@input_file) do
      {:ok, body} ->
        checksum2(mapify(body))
      {:error, reason} ->
        {:error, reason}
    end
  end
end
