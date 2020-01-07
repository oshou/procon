def minimize(arr1, arr2, sum)
  p arr1
  p arr2
  p sum
  if (arr1.empty?) && (arr2.empty?)
    return sum
  end

  min = nil
  minv = 0
  arr1.product(arr2).each do |v|
    d = (v[0] - v[1]).abs
    if min == nil
      min = d
      minv = v
    end

    if min > d
      min = d
      minv = v
    end
  end
  sum += min
  arr1.delete(minv[0])
  arr2.delete(minv[1])
  minimize(arr1, arr2, sum)
end

# Input
nums = []
3.times { |i| nums += gets.chomp.split.map(&:to_i) }

# Counts
counts = (1..9).to_a.map { |i| [i, 0] }.to_h
nums.each do |num|
  counts[num] += 1
end

duplicates = counts.select { |k, v| v > 1 }.keys
needs = counts.select { |k, v| v == 0 }.keys
puts minimize(duplicates, needs, 0)
