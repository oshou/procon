a, b, c, d = gets.chomp.split.map { |i| i.to_i }
puts (a * d - b * c).abs / 2
