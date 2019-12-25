n, m = gets.chomp.split.map { |i| i.to_i }
ans = n+m
if ans < 16
  puts "HIT"
else
  puts "STAND"
end
