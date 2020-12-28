n = int(input())
arr = list(map(int, input().split()))
cnt = 0
for i in range(n):
    for j in range(i+1, n):
        for k in range(j+1, n):
            if (arr[i] < arr[j]+arr[k]) and (arr[j] < arr[k]+arr[i]) and (arr[k] < arr[i]+arr[j]):
                cnt += 1
print(cnt)
