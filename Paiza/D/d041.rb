n, d, e = gets.split.map(&:to_i)
if n <= d * e
  puts "OK"
else
  puts "NG"
end
