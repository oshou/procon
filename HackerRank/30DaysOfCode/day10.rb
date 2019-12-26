str = gets.to_i.to_s(2)
token = ""
tokens = []
str.each_char do |char|
  if (token == "") || (token[-1] == char)
    token += char
  else
    tokens << token
    token = char
  end
end
tokens << token
puts tokens.max.length
