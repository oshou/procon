m, n = gets.split.map(&:to_i)
for i in 0..9
  if i == 0
    printf("%d", m + i * n)
  else
    printf(" %d", m + i * n)
  end
end
print "\n"
