A = list(map(int,input().split()))

num = 0
for i in A:
    for j in A:
        num += A[i] - A[j]

print(num)
