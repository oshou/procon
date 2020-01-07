arr = gets.chomp.split("")
puniaki = "Yes"
arr.each_with_index do |char, i|
  if (i % 2 != 0)
    if char != " "
      puniaki = "No"
      break
    end
  else
    if !(char.match?(/[a-z]/))
      puniaki = "No"
      break
    end
  end
end
puts puniaki
