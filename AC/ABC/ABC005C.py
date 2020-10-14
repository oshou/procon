T = int(input())
N = int(input())
A = list(map(int, input().split()))
M = int(input())
B = list(map(int, input().split()))

b_asc = sorted(B)
a_asc = sorted(A)

ans = "yes"
for i in range(M):
    diff = b_asc[i] - a_asc[i%N]
    if (diff < 0) or (T < diff):
        ans = "no"

print(ans)
