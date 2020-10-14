N, A, B = map(int, input().split())

ans = 0
for i in range(N+1):
    digit_sum = sum(list(map(int, (list(str(i))))))
    if A <= digit_sum <= B:
        ans += digit_sum

print(ans)
