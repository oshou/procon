N, A, B = map(int, input().split())

total = 0
for i in range(1, N+1):
    S = sum([int(x) for x in str(i)])
    if A <= S <= B:
        total += i

print(total)
