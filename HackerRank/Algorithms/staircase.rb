#!/bin/ruby

require "json"
require "stringio"

# Complete the staircase function below.
def staircase(n)
  s = " " * n
  (1..n).each do |i|
    s[-i] = "#"
    puts s
  end
end

n = gets.to_i
staircase(n)
