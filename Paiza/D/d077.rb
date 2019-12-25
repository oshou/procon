a, b = gets.chomp.split.map { |i| i.to_i }
ans = a*b
if ans > 10000
  puts "NG"
else
  puts ans
end
