n, k = map(int, input().split())
h = list(map(int, input().split()))
h.sort()
sum = 0
if n <= k:
    print(0)
    exit()
for i in range(n-k):
    sum += h[i]
print(sum)
