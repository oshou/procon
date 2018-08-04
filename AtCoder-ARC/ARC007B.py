N, M = map(int, input().split())
case = [ i for i in range(N+1)]

for _ in range(M):
    disk = int(input())
    index = case.index(disk)
    case[index] = case[0]
    case[0] = disk

for i in range(1, N+1):
    print(case[i])
