n = gets.to_i
str = gets.chomp
pos = 0
cnt = 0
str.chars.each do |char|
  if char == "U"
    if pos == (-1)
      cnt += 1
    end
    pos += 1
  else
    pos -= 1
  end
end
puts cnt
