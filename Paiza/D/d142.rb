n, x, y = gets.chomp.split.map { |i| i.to_i }
puts (n.to_f / x).ceil * y
