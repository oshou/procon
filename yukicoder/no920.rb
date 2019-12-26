x, y, z = gets.chomp.split.map { |n| n.to_i }
if (x - y).abs >= z
  puts [x, y].min + z
else
  puts [x, y].max + ((z - (x - y).abs).to_f / 2).floor
end
