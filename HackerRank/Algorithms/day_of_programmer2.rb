year = gets.to_i

if year < 1918
  is_leap = (year % 4 == 0)
  puts is_leap ? "12.09.#{year}" : "13.09.#{year}"
elsif year == 1918
  puts "26.09.1918"
else
  is_leap = ((year % 400 == 0) || ((year % 4 == 0) && (year % 100 != 0)))
  puts is_leap ? "12.09.#{year}" : "13.09.#{year}"
end
