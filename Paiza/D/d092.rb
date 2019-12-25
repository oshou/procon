x1, y1, p1 = gets.split.map { |i| i.to_i }
x2, y2, p2 = gets.split.map { |i| i.to_i }
c1 = p1.to_f / (x1 * y1)
c2 = p2.to_f / (x2 * y2)
if c1 == c2
  puts "DRAW"
elsif c1 < c2
  puts "#{x1} #{y1} #{p1}"
else
  puts "#{x2} #{y2} #{p2}"
end
