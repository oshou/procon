n, k = gets.chomp.split.map(&:to_i)
arr = gets.chomp.split.map(&:to_i)

cnt = 0
(0..n - 1).each do |i|
  (0..n - 1).each do |j|
    if (i < j) && ((arr[i] + arr[j]) % k == 0)
      cnt += 1
    end
  end
end
puts cnt
