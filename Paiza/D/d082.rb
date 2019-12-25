a = gets.chomp
b = gets.chomp
if (a[-1] == b[0]) && (b[-1] != "n")
  puts "OK"
else
  puts "NG"
end
