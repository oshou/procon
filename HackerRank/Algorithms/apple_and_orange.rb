s, t = gets.chomp.split.map(&:to_i)
a, b = gets.chomp.split.map(&:to_i)
m, n = gets.chomp.split.map(&:to_i)
distsa = gets.chomp.split.map(&:to_i)
distsb = gets.chomp.split.map(&:to_i)

cnta = 0
distsa.each { |dist|
  pos = a + dist
  if (s..t).include?(pos)
    cnta += 1
  end
}

cntb = 0
distsb.each { |dist|
  pos = b + dist
  if (s..t).include?(pos)
    cntb += 1
  end
}

puts cnta
puts cntb
