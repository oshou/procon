arr = gets.chomp.split.map { |i| i.to_i }
cloudy = arr.select { |i| i == 1 }.length
if cloudy >= 5
  puts "yes"
else
  puts "no"
end
