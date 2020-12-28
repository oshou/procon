n = int(input())
mmin = n-1
for i in range(1, n+1):
    if i*i > n:
        break
    if n % i == 0:
        mmin = min(mmin, (i-1)+(n//i-1))
print(mmin)
