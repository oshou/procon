n, a, b = map(int, input().split())
sum = 0
for i in range(1, n+1):
    dsum = 0
    tmp = i
    while tmp != 0:
        dsum += tmp % 10
        tmp //= 10
    if a <= dsum and dsum <= b:
        sum += i
print(sum)
