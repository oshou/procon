n = gets.chomp.to_i
phonebook = {}
n.times {
  name, number = gets.chomp.split
  phonebook[name] = number
}

n.times do
  name = gets.chomp
  if phonebook.has_key?(name)
    puts "#{name}=#{phonebook[name]}"
  else
    puts "Not found"
  end
end
