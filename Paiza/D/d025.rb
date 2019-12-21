n = gets.to_i
s = n.to_s
case s.length
when 3
  puts s
when 2
  puts "0" + s
when 1
  puts "00" + s
else
  puts "000"
end
