a, b = gets.chomp.split.map { |i| i.to_i }
puts (b * (1 + (a.to_f / 100))).floor
