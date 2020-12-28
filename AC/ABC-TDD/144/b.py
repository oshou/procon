n = int(input())
for i in range(1, n+1):
    if i > 9:
        print("No")
        exit()
    if i*i > n:
        print("No")
        exit()
    res = n//i
    if n % i == 0 and 1 <= res and res <= 9:
        print("Yes")
        exit()
print("No")
