h, m = gets.split(":").map { |i| i.to_i }
if h >= 8
  puts "#{h - 8}:#{m}"
else
  puts "#{h + 16}:#{m}"
end
