N, K = map(int, input().split())
R = list(map(int, input().split()))

R.sort()
partial = R[-K:]
ans = 0
for i in partial:
    ans = (ans+i)/2

print(ans)
