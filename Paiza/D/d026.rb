arr = 7.times.map { gets.chomp }
cnt = 0
arr.each { |item|
  if item == "no"
    cnt += 1
  end
}
puts cnt
