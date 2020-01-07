n = gets.to_i
arr = gets.chomp.split.map(&:to_i)

counts = {}
uniques = arr.uniq
uniques.each do |v|
  counts[v] = arr.count(v)
end

puts counts.sort.max { |k, v| k[1] <=> v[1] }[0]
