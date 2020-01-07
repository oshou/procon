#!/bin/ruby

require "json"
require "stringio"

# Complete the miniMaxSum function below.
def miniMaxSum(arr)
  sums = []
  (0..arr.length - 1).each do |i|
    tmp = arr.dup
    tmp.delete_at(i)
    sums[i] = tmp.sum
  end
  puts "#{sums.min} #{sums.max}"
end

arr = gets.rstrip.split(" ").map(&:to_i)

miniMaxSum arr
