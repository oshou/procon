n = gets.to_i
(0...n).each do |i|
  puts " " * (n - i - 1) + "#" * (i + 1)
end
