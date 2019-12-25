w, h = gets.chomp.split.map { |i| i.to_i }
arr = h.times.map { gets.chomp }
arr.each do |str|
  puts str
end
