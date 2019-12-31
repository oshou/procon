n = gets.to_i
n.times {
  catA, catB, mouseC = gets.chomp.split.map(&:to_i)
  distA = (mouseC - catA).abs
  distB = (mouseC - catB).abs

  if distA == distB
    puts "Mouse C"
  elsif distA < distB
    puts "Cat A"
  else
    puts "Cat B"
  end
}
