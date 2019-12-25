n = gets.to_i
arr = n.times.map { gets.to_i }
puts (arr.sum.to_f / arr.length).floor
