medals = 3.times.map { gets.chomp }
medals.each_with_index { |value, i|
  case i
  when 0
    puts "Gold #{value}"
  when 1
    puts "Silver #{value}"
  when 2
    puts "Bronze #{value}"
  end
}
