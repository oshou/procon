#a, b, k = map(int, input().split())
#maxint = max(a, b)
#cnt = 0
# for i in range(maxint, 0, -1):
#    if a % i == 0 and b % i == 0:
#        cnt += 1
#    if cnt == k:
#        print(i)
#        break
n = int(input())
prime_cnt = 0
cnt = 0
for i in range(1, n+1):
    for j in range(1, n+1):
        if i % j == 0:
            prime_cnt += 1
    if prime_cnt == 8:
        cnt += 1
print(cnt)
