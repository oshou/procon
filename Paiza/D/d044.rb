s1, s2 = gets.split.map { |str| str.chomp }
case s2
when "M"
  puts "Hi, Mr.#{s1}"
when "F"
  puts "Hi, Ms.#{s1}"
end
