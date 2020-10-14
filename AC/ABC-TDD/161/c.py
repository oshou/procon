import math

n, k = map(int, input().split())
m = n/k
ans1 = abs(n - math.floor(m)*k)
ans2 = abs(n - math.ceil(m)*k)
print(min(ans1, ans2))
