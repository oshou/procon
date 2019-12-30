#!/bin/ruby

require "json"
require "stringio"

# Complete the plusMinus function below.
def plusMinus(arr)
  plus = []
  minus = []
  zero = []
  arr.each do |v|
    if v == 0
      zero << v
    elsif v > 0
      plus << v
    else
      minus << v
    end
  end
  total_count = arr.length
  puts sprintf("%.6f", plus.length.to_f / total_count)
  puts sprintf("%.6f", minus.length.to_f / total_count)
  puts sprintf("%.6f", zero.length.to_f / total_count)
end

n = gets.to_i
arr = gets.rstrip.split(" ").map(&:to_i)

plusMinus arr
