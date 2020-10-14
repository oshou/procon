n = gets.to_i
arr = n.times.map { gets.chomp }

ans = "DRAW"
arr.each_index do |i|
  if i == 0
    next
  end
  ng = ((arr[i][0] != arr[i - 1][-1]) || (arr.count(arr[i]) >= 2))
  if ng
    if (i % 2 == 0)
      ans = "WIN"
    else
      ans = "LOSE"
    end
  end
end
puts ans
