c = []
4.times { |i|
  c[i] = gets.chomp.split
}
p c

c2 = [[], [], [], []]
(0..3).each { |i|
  (0..3).each { |j|
    c2[-(i + 1)][-(j + 1)] = c[i][j]
  }
}

(0..3).each { |i|
  puts c2[i].join(" ")
}
