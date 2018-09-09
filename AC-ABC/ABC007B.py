A = str(input())

ans = -1
for i in range(len(A)):
    if A[i] > "a":
        ans = "a" * (i + 1)
        break

if len(A) > 1:
    ans = "a" * (len(A) - 1)

print(ans)
