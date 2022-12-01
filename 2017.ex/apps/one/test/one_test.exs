defmodule OneTest do
  use ExUnit.Case

  test "captcha/1 return correct values" do
    assert One.captcha(1122) == 3
    assert One.captcha(1111) == 4
    assert One.captcha(1234) == 0
    assert One.captcha(91212129) == 9
  end

  # second implementation
  test "catpcha2/1 returns correct values" do
    assert One.captcha2(1122) == 3
    assert One.captcha2(1111) == 4
    assert One.captcha2(1234) == 0
    assert One.captcha2(91212129) == 9
  end

  test "halve/1 cuts an even sized array in half" do
    assert One.halve([1, 2, 3, 4]) == [[1, 2], [3, 4]]
  end
  
  # part two
  test "captcha3/1 returns correct values" do
    assert One.captcha3(1212) == 6
    assert One.captcha3(1221) == 0
    assert One.captcha3(123425) == 4
    assert One.captcha3(123123) == 12
    assert One.captcha3(12131415) == 4
  end
end
