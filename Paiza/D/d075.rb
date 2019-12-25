arr = [1, 2, 3, 4, 5]
4.times {
  n = gets.to_i
  arr.delete(n)
}
puts arr[0]
