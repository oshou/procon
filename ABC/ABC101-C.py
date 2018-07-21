import math

n, k = map(int, input().split())
arr = list(map(int, input().split()))

if n == k:
    print(1)
    exit()
else:
    print(1 + math.ceil((n-k)/(k-1)))
