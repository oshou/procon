n = gets.to_i
arr = gets.chomp.split.map(&:to_i)
d, m = gets.chomp.split.map(&:to_i)
cnt = 0
(0..arr.length - m).each do |i|
  if arr.slice(i, m).sum == d
    cnt += 1
  end
end
puts cnt
