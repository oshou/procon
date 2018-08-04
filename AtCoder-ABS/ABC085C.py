N, Y = map(int, input().split())

for i in range(N+1):
    for j in range(N+1):
        for k in range(N+1):
            if (i + j + k == N) and (10000*i + 5000*j + 1000*k == Y):
                print(i, j, k)
                exit()
print(-1, -1, -1)
