budget, n, m = gets.chomp.split.map(&:to_i)
keyboards = gets.chomp.split.map(&:to_i)
drives = gets.chomp.split.map(&:to_i)

max = -1
keyboards.each do |k|
  drives.each do |d|
    paid = k + d
    if (paid <= budget) && (max < paid)
      max = paid
    end
  end
end
puts max
