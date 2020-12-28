n = int(input())
arr = list(map(int, input().split()))
maxdiv, maxcnt = 0, 0
MAX = max(32, max(arr)+1)
for i in range(2, MAX):
    cnt = 0
    for j in range(n):
        if arr[j] % i == 0:
            cnt += 1
    if maxcnt < cnt:
        maxcnt = cnt
        maxdiv = i
print(maxdiv)
