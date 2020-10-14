D, G = map(int, input().split())
arr = [[]]
base_ef = [0]
total_ef = [0]
for i in range(1, D+1):
    arr.append(list(map(int, input().split())))
    base_ef.append(100*i)
    total_ef.append(100*i + arr[i][1]/arr[i][0])

cnt = 0
rem = G
while rem =< 0:
    rem = rem - max(total_ef)
    cnt += 1

print(cnt)
