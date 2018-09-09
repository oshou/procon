N = int(input())
arr = [input() for i in range(N)]
ans = "Yes"

if len(set(arr)) != N:
    ans = "No"

for i in range(N):
    if not i == 0:
        if arr[i-1][-1] != arr[i][0]:
            ans = "No"
            break

print(ans)
