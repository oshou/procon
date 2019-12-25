arr = 3.times.map { gets.chomp }
d = 0
c = 0
arr.each { |s|
  if s == "cat"
    c += 1
  else
    d += 1
  end
}
if c > d
  puts "cat"
else
  puts "dog"
end
