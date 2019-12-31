n = gets.to_i
arr = gets.chomp.split.map(&:to_i)
tallest = arr.max
puts arr.count(tallest)
