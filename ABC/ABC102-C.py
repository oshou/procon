n = int(input())
A = list(map(int, input().split()))

SAN = []
sum = 0
for b in range(1, n+1):
    for i in range(1, n+1):
        print(i)
        sum += abs(A[i]-(b+i))
        print(sum)
    SAN.append(sum)

print(min(SAN))
