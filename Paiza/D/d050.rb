reqs = gets.split.map { |n| n.to_i }
sum = 0
reqs.each { |v|
  if v >= 5
    sum += 5
  else
    sum += v
  end
}
puts sum
