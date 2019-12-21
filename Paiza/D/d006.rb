n, str = gets.chomp.split()
n = n.to_i
case str
when "km"
  puts n * 1000 * 100 * 10
when "m"
  puts n * 1000 * 10
when "cm"
  puts n * 10
end
