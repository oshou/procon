x, y, p = gets.chomp.split.map { |i| i.to_i }
puts p * (x.to_f / y).ceil
