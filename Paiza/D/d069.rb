arr = gets.chomp.split.map { |i| i.to_i }
puts (arr.sum.to_f / arr.length).round(1)
