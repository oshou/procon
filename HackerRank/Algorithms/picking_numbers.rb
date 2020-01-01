n = gets.chomp.to_i
arr = gets.chomp.split.map(&:to_i)

counts = {}
arr.uniq.each do |v|
  counts[v] = arr.count(v)
end

product = counts.keys.product(counts.keys)

max = 0
product.each do |v|
  if (v[0] - v[1]).abs > 1
    next
  end

  if (v[0] == v[1])
    sum = counts[v[0]]
  else
    sum = counts[v[0]] + counts[v[1]]
  end

  if (max < sum)
    max = sum
  end
end
puts max
