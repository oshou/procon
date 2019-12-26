def factorial(n)
  if n <= 1
    return 1
  else
    return n * factorial(n - 1)
  end
end

n = gets.to_i
puts factorial(n)
