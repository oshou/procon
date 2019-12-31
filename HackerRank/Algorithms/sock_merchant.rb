n = gets.chomp.to_i
arr = gets.chomp.split.map(&:to_i)

counts = {}
arr.uniq.each do |v|
  counts[v] = arr.count(v)
end

pairs = 0
counts.map { |k, v| pairs += v / 2 }
puts pairs
