n = gets.to_i
arr = gets.chomp.split
capital = ""
arr.each do |i|
  capital += i[0]
end
puts capital
