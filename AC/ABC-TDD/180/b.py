import math
n = int(input())
x = list(map(int, input().split()))
m, us, c = 0, 0, 0
for i in range(n):
    m += abs(x[i])
    us += x[i]*x[i]
    c = max(abs(x[i]), c)
u = math.sqrt(us)
print(m)
print(u)
print(c)
