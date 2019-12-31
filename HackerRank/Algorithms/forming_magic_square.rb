def minimize(arr1, arr2, sum)
  if (arr1.empty?) && (arr2.empty?)
    return sum
  end

  min = nil
  minv = 0
  arr1.product(arr2).each do |v|
    d = (v[0] - v[1]).abs
    if min == nil
      min = d
    end

    if min > d
      min = d
      minv = v
    end
  end
  sum += min
  arr1.delete_at(minv[0])
  arr2.delete_at(minv[1])
  minimize(arr1, arr2, sum)
end

nums = []
3.times { |i|
  nums += gets.chomp.split.map(&:to_i)
}

counts = {
  1 => 0,
  2 => 0,
  3 => 0,
  4 => 0,
  5 => 0,
  6 => 0,
  7 => 0,
  8 => 0,
  9 => 0,
}

nums.each do |num|
  counts[num] += 1
end
duplicates = counts.select { |k, v| v > 1 }.keys
needs = counts.select { |k, v| v == 0 }.keys
puts minimize(duplicates, needs, 0)
