defmodule Four do
  @moduledoc """
  advent of code, day 4: high entropy passphrases

  """

  @input_file "four.txt"

  def read_input(f) do
    case File.read(f) do
      {:ok, body} ->
        body
      {:error, reason} ->
        {:error, reason}
    end
  end

  def solve_one do
    read_input(@input_file)
    |> String.split("\n", trim: :true) 
    |> Enum.reduce(0, fn(cur, acc) -> if (validate(cur)), do: acc + 1, else: acc end)
  end

  def word_count([head | tail]) do
    word_count(tail, %{head => 1})
  end

  def word_count([head | tail], counter) do
    {_, new_counter} = Map.get_and_update(counter, head, fn current -> 
      case current do
        nil -> {current, 1}
        _   -> {current, current + 1}
      end
    end)
    word_count(tail, new_counter)
  end

  def word_count([], counter) do
    counter
  end

  def validate(str) do
    case str
          |> String.split(" ")
          |> word_count
          |> Map.values
          |> Enum.filter(fn x -> x > 1 end)
          |> length
    do
      0 -> :true
      _ -> :false
    end
  end

  # part two

  def sort_word(word) do
    word |> String.graphemes |> Enum.sort |> Enum.join
  end

  def contains_anagram_pair?(arr) do
    arr 
    |> Enum.map(&sort_word/1)
    |> Enum.uniq
    |> length !== length(arr)
  end

  def solve_two do
    read_input(@input_file)
    |> String.split("\n", trim: :true) 
    |> Enum.map(fn el -> String.split(el, " ") end)
    |> Enum.reduce(0, fn(cur, acc) -> if (contains_anagram_pair?(cur)), do: acc, else: acc + 1 end)
  end
end

