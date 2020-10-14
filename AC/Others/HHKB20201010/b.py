h, w = map(int, input().split())
s = [""]*h
cnt = 0
for i in range(0, h):
    s[i] = input()
for i in range(0, h):
    for j in range(0, w-1):
        if s[i][j] == "." and s[i][j+1] == ".":
            cnt += 1

for i in range(0, h-1):
    for j in range(0, w):
        if s[i][j] == "." and s[i+1][j] == ".":
            cnt += 1

print(cnt)
