arr = gets.chomp.split(",")
if arr.select { |str| str != "AC" }.empty?
  puts "Done!"
else
  puts "Failed..."
end
