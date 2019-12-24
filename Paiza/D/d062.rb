str = "ABCDEFGHIJ"
h = gets.split.map { |i| i.to_i }
puts str.slice!(0, h[0])
puts str.slice!(0, h[1])
puts str.slice!(0, h[2])
