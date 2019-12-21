str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
s = gets.chomp
for i in 0..str.length
  if str[i] == s
    puts i + 1
    exit
  end
end
