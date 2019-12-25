p = gets.split.map { |i| i.to_i }
t = 0
p.each { |n|
  if n <= 2
    t += 1
  end
}
puts t
