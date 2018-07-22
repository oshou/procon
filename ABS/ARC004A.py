import math

N = int(input())
x = []
y = []
for _ in range(N):
    xi, yi = map(int, input().split())
    x.append(int(xi))
    y.append(int(yi))

l = 0
for i in range(N):
    for j in range(i+1, N):
        d = (x[i] - x[j]) * (x[i] - x[j]) + (y[i] - y[j]) * (y[i] - y[j])
        l = max(l, d)

print sqrt(l)
