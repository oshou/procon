t=gets.to_i
arr = t.times.map { gets.chomp }
arr.each do |str|
  s1 = ""
  s2 = ""
  str.each_char.with_index{ |v,i|
    if i % 2 == 0
      s1 << v
    else
      s2 << v
    end
  }
  puts "#{s1} #{s2}"
end
