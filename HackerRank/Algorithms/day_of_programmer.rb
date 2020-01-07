year = gets.to_i

case year
when 1700..1917
  if year % 4 == 0
    puts "12.09.#{year}"
  else
    puts "13.09.#{year}"
  end
when 1918
  puts "26.09.#{year}"
else
  if (year % 400 == 0) || ((year % 4 == 0) && (year % 100 != 0))
    puts "12.09.#{year}"
  else
    puts "13.09.#{year}"
  end
end
