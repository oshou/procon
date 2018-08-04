K, S = map(int, input().split())

cnt = 0
for i in range(K+1):
    for j in range(K+1):
        for k in range(K+1):
            if i + j + k == S:
                cnt += 1

print(cnt)
