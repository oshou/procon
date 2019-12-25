n = gets.to_i
h, w = gets.chomp.split.map { |i| i.to_i }
puts h * w % n
