N, X = map(int, input().split())
m = [ int(input()) for _ in range(N) ]
ans = (X - sum(m)) // min(m) + N

print(ans)
