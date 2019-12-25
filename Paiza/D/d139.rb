n = gets.to_i
arr = gets.chomp.split
g = arr.count("G")
p = arr.count("P")
if g == p
  puts "Draw"
elsif g > p
  puts "P"
else
  puts "G"
end
