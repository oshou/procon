arr = gets.chomp.split.map { |i| i.to_i }
x = gets.to_i
avg = arr.sum.to_f / arr.length
if avg <= x
  puts "pass"
else
  puts "failure"
end
