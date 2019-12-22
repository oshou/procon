days = 5.times.map { |day| gets.to_i }
days.each_index { |idx|
  if idx > 0
    puts days[idx] - days[idx - 1]
  end
}
