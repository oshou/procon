require "date"

s = gets.chomp

hour = s[0..1].to_i
ap = s.slice!(s.length - 2, 2)
case ap
when "AM"
  if hour >= 12
    s[0..1] = sprintf("%02d", hour - 12)
  end
when "PM"
  if hour < 12
    s[0..1] = sprintf("%02d", hour + 12)
  end
end
puts s
