n, k = map(int, input().split())
sunuke = [False]*n

for i in range(k):
    d = int(input())
    arr = list(map(int, input().split()))
    for j in range(d):
        sunuke[arr[j]-1] = True

cnt = 0
for v in sunuke:
    if not v:
        cnt += 1
print(cnt)
