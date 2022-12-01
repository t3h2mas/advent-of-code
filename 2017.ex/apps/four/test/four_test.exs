defmodule FourTest do
  use ExUnit.Case

  test "validates a passphrase" do
    assert Four.validate("aa bb cc dd ee") == :true
  end

  test "does not allow duplicate words" do
    assert Four.validate("aa bb cc dd aa") == :false
  end

  test "allows aa and aaa" do
    assert Four.validate("aa bb cc dd aaa") == :true
  end

  test "word_count can count words" do
    assert Four.word_count(["hi"]) == %{"hi" => 1}
    assert Four.word_count(["hi", "hi"]) == %{"hi" => 2}
    assert Four.word_count(["hi", "hi", "bye"]) == %{"bye" => 1, "hi" => 2}
  end

  test "sort_word/1 sorts a word by letters" do
    assert Four.sort_word("teddy") == "ddety"
  end

  test "contains_anagram_pair?/1 detects anagrams" do
    assert Four.contains_anagram_pair?(["one", "eno"]) == :true
    assert Four.contains_anagram_pair?(["onee", "eno"]) == :false
  end
end
