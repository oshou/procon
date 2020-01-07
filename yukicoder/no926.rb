a, b, c = gets.chomp.split.map { |i| i.to_i }
puts sprintf("%.7f", ((a.to_f / b) * c).to_f)
