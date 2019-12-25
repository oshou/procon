str = gets.chomp
hash = {
  0 => "C",
  1 => "A",
  2 => "B",
}
m = ""
str.each_char do |s|
  m += hash[s.to_i]
end
puts m
