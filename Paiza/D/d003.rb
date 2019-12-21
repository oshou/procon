n = gets.to_i
for i in 1..9
  if i == 1
    printf("%d", i * n)
  else
    printf(" %d", i * n)
  end
end
