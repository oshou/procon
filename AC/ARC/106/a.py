n = int(input())
for i in range(1, 38):
    a = 3**i
    for j in range(1, 26):
        b = 5**j
        if n == a + b:
            print(i, j)
            exit()
print(-1)
