import math

N, X = map(int, input().split())
m = [int(input()) for _ in range(N)]
cnt = N + math.floor((X - sum(m)) / min(m))
print(cnt)
