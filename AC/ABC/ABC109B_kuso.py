N = int(input())
ans = "Yes"
arr = []
for i in range(N):
    arr.append(input())
    if not i == 0:
        if arr[i-1][-1] == arr[i][0] and not arr[i] in arr[:-2]:
            pass
        else:
            ans = "No"
            break

print(ans)
