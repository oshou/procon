def solve(n)
  if n % 2 == 1
    return "Weird"
  elsif (2..5).include?(n)
    return "Not Weird"
  elsif (6..20).include?(n)
    return "Weird"
  else
    return "Not Weird"
  end
end

n = gets.to_i
puts solve(n)

require "minitest/autorun"

class HackerRankTest < Minitest::Test
  def test_solve
    assert_equal("Weird", solve(1))
    assert_equal("Weird", solve(3))
    assert_equal("Weird", solve(13))
    assert_equal("Not Weird", solve(2))
    assert_equal("Not Weird", solve(4))
    assert_equal("Weird", solve(6))
    assert_equal("Weird", solve(12))
    assert_equal("Weird", solve(16))
    assert_equal("Not Weird", solve(22))
  end
end
