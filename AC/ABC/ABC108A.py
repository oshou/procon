K = int(input())

even = K // 2

if (K % 2):
    odd = K // 2 + 1
else:
    odd = K // 2

ans = even * odd
print(ans)
