n, m = gets.chomp.split.map(&:to_i)
arra = gets.chomp.split.map(&:to_i)
arrb = gets.chomp.split.map(&:to_i)

cnt = 0
(arra.min..arrb.min).each do |v|
  if (arra.select { |n| v % n == 0 } == arra) && (arrb.select { |n| n % v == 0 } == arrb)
    cnt += 1
  end
end
puts cnt
