n = gets.chomp.to_i
n.times {
  cnt = 0
  num = gets.chomp
  num.chars.each { |char|
    if (char != "0") && (num.to_i % char.to_i == 0)
      cnt += 1
    end
  }
  puts cnt
}
