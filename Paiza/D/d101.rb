n, m = gets.split.map { |i| i.to_i }
is_odd_n = (n % 2) > 0
is_odd_m = (m % 2) > 0
if (is_odd_n && is_odd_m) || (!is_odd_n && !is_odd_m)
  puts "NO"
else
  puts "YES"
end
