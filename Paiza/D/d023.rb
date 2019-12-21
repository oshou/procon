str = gets.chomp
cnt = 0
str.each_char do |char|
  if char == "A"
    cnt += 1
  end
end
puts cnt
