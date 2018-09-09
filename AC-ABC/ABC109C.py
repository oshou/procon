def gcd(a, b):
    if b == 0:
        return a
    else:
        return gcd(b, a%b)

N, start = map(int, input().split())
arr = list(map(int, input().split()))

arr.append(X)
arr.sort()
diff = []
for i in range(N):
    diff.append(abs(arr[i+1] - arr[i]))

ans = max(diff)
for d in diff:
    ans = gcd(g, d)

print(ans)
