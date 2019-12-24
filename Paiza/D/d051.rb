cnt = 0
arr = gets.split.map { |str|
  if str.chomp == "W"
    cnt += 1
  end
}

if cnt >= 5
  puts "OK"
else
  puts "NG"
end
