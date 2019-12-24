t = gets.split.map { |i| i.to_i }
a = gets.to_i
num = 0
t.each { |n|
  puts n
  if n > a
    exit
  else
    num += 1
  end
}
puts num
