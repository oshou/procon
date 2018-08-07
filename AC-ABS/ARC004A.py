import math

N = int(input())
arr = [ list(map(int, input().split())) for _ in range(N) ]

max_length = 0
for i in range(N):
    for j in range(N):
        max_length = max(max_length, math.sqrt((arr[j][0] - arr[i][0]) ** 2 + (arr[j][1] - arr[i][1]) ** 2))

print(max_length)
