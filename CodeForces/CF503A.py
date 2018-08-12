# n = voters
# m = parties
n, m = map(int, input().split())
p = [ 0 for _ in range(n)]
c = [ 0 for _ in range(n)]
for i in range(n):
    p[i], c[i] = map(int, input().split())

max = 0
for i in range(n):
    if p.count(i) >= max:
        max = p.count(i)

for i in range(M):

ans = 0
while True:
    if ans > max:
        print(ans)
        break;
    else:
        ans += points[i]
        i += 1
