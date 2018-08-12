N = int(input())
ans = "No"
for i in range(25):
    for j in range(15):
        if 4 * i + 7 * j == N:
            ans = "Yes"

print(ans)
