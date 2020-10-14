m, n, N = map(int, input().split())

sum = target = reuse = N
while reuse > 0:
    reuse = target // m * n
    unused = target % m
    sum += reuse
    target = reuse + unused

print(sum)
