m, n = gets.split.map { |i| i.to_i }
remain = n-m
if remain >= 0
  puts remain
else
  puts "No"
end
