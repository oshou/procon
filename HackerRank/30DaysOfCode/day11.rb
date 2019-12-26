arr = []
(0..5).each { |i|
  arr[i] = gets.chomp.split.map { |n| n.to_i }
}

max = -63
(0..3).each { |i|
  (0..3).each { |j|
    h = arr[i][j] + arr[i][j + 1] + arr[i][j + 2] + arr[i + 1][j + 1] + arr[i + 2][j] + arr[i + 2][j + 1] + arr[i + 2][j + 2]
    max = [h, max].max
  }
}
puts max
