grade = gets.to_i
presents = gets.chomp.split
case grade
when 3
  puts presents[-1]
when 2
  puts presents[-2]
when 1
  puts presents[-3]
end
