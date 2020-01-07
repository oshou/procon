x1, v1, x2, v2 = gets.chomp.split.map(&:to_i)
if v1 == v2
  if x1 == x2
    puts "YES"
  else
    puts "NO"
  end
  exit
end

n = (x2 - x1) / (v1 - v2)
m = (x2 - x1) % (v1 - v2)
if (n >= 0) && (m == 0)
  puts "YES"
else
  puts "NO"
end
