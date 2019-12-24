str = gets.chomp
if str.length >= 11
  puts "OK"
else
  puts 11 - str.length.to_i
end
