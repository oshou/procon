n = int(input())
s = input()
if n % 2 == 1:
    print("No")
    exit()
mid = n//2
for i in range(0, mid):
    if s[i] != s[i+mid]:
        print("No")
        exit()
print("Yes")
