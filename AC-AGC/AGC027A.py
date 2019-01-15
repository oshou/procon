N, x = map(int, input().split())
arr = list(map(int, input().split()))

arr.sort()
cnt = 0
for ele in arr:
    if x < ele:
        break
    x -= ele
    cnt += 1

if cnt == N and x > 0:
    cnt -= 1

print(cnt)
