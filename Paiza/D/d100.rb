s = gets.chomp
if s.count("_") >= s.count("-")
  puts s.gsub(/-/, "_")
else
  puts s.gsub(/_/, "-")
end
