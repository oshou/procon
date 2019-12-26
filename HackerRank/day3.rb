n=gets.to_i
if n % 2 != 0
  puts "Weird"
else
  if (2..5).include?(n)
    puts "Not Weird"
  elsif (6..20).include?(n)
    puts "Weird"
  else
    puts "Not Weird"
  end
end
