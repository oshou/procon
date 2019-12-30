#!/bin/ruby

require "json"
require "stringio"

def diagonalDifference(arr)
  ans = 0
  (0..arr.length - 1).each do |i|
    ans += arr[i][i] - arr[i][-i - 1]
  end
  return ans.abs
end

fptr = File.open(ENV["OUTPUT_PATH"], "w")

n = gets.chomp.to_i
arr = Array.new(n)
n.times do |i|
  arr[i] = gets.chomp.split.map(&:to_i)
end

result = diagonalDifference(arr)
p result

fptr.write result
fptr.write "\n"

fptr.close()
