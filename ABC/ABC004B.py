N = 4
arr = []
for i in range(N):
    row = list(map(str, input().split()))
    arr.append(row)

for i in range(N):
    for j in range(N):
        print(arr[-(i+1)][j], end="")
        if j < N-1:
            print(" ", end="")
        else:
            print("")
