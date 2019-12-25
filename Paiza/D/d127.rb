s = gets.chomp
s.slice!(0)
if s[0] * s.length == s
  puts "Yes"
else
  puts "No"
end
