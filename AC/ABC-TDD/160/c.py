k, n = map(int, input().split())
a = list(map(int, input().split()))
mid = k-a[n-1]+a[0]
for i in range(0, n):
    if i == 0:
        continue
    tmp = a[i]-a[i-1]
    if mid < tmp:
        mid = tmp
print(k-mid)
