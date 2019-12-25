str = gets.chomp
ans = str
str.each_char { |s|
  if (s == "I" || s == "l" || s == "i")
    ans= "caution"
  end
}
puts ans
