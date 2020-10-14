n = int(input())
a = 0
b = 0
arr = list(map(int, input().split()))
arr.sort(reverse=True)
for i in range(0, n):
    if i % 2 == 0:
        a += arr[i]
    else:
        b += arr[i]

print(a-b)
