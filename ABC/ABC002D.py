import math

N, M = map(int, input().split())

arr = [[ False for _ in range(N+1)] for _ in range(N+1)]
for i in range(N):
    arr[i][i] = True

for _ in range(M):
    x, y = map(int, input().split())
    arr[x][y] = True
    arr[y][x] = True

ans = 0
for i in range(N):
    if ans < arr[i].count(True):
        ans = arr[i].count(True)

print(ans)
