s = gets.chomp
s.each_char.with_index { |str, i|
  print str
  if (i + 1) % 10 == 0
    print "\n"
  end
}
