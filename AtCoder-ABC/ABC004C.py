N = int(input())
arr = [1, 2, 3, 4, 5, 6]

for i in range(N):
    arr[i%5], arr[i%5+1] = arr[i%5+1], arr[i%5]

for i in range(6):
    print(arr[i], end="")
