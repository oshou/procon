n, m = map(int, input().split())
a = list(map(int, input().split()))
border = sum(a)/(4*m)
cnt = 0
for i in range(n):
    if a[i] >= border:
        cnt += 1
if cnt >= m:
    print("Yes")
else:
    print("No")
