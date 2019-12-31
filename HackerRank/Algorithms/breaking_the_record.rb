n = gets.to_i
arr = gets.chomp.split.map(&:to_i)

min = max = nil
min_count = max_count = 0
arr.each_with_index do |v, i|
  if i == 0
    min = max = v
    next
  end

  if v > max
    max = v
    max_count += 1
  end

  if v < min
    min = v
    min_count += 1
  end
end
puts "#{max_count} #{min_count}"
