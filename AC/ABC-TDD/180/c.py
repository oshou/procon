n = int(input())
ans = []
for i in range(1, n+1):
    if i*i > n:
        break
    if n % i == 0:
        ans.append(i)
        tmp = n//i
        if i != tmp:
            ans.append(n//i)
ans = sorted(ans)
counts = len(ans)
for num in ans:
    print(num)
