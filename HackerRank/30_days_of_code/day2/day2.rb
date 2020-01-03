#!/bin/ruby

require "json"
require "stringio"

# Complete the solve function below.
def solve(meal_cost, tip_percent, tax_percent)
  tip_cost = meal_cost * (tip_percent / 100.0)
  tax_cost = meal_cost * (tax_percent / 100.0)
  return (meal_cost + tip_cost + tax_cost).round
end

#meal_cost = gets.to_f
#tip_percent = gets.to_i
#tax_percent = gets.to_i
#solve meal_cost, tip_percent, tax_percent

require "minitest/autorun"

class HackerRankTest < MiniTest::Test
  def test_solve
    assert_equal(15, solve(12.00, 20, 8))
    assert_equal(19, solve(15.50, 15, 10))
  end
end
