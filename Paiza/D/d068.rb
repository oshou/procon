n = gets.to_i
s = gets.chomp
strS = 0
strR = 0
(0..n - 1).each do |i|
  if s[i] == "S"
    strS += 1
  else
    strR += 1
  end
end
puts "#{strS} #{strR}"
