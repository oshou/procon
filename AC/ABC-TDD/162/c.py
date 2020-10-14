def gcd(a, b: int) -> int:
    if a < b:
        a, b = b, a
    if b == 0:
        print("a,b,ans:", a, b, a)
        return a
    print("a,b,ans:", a, b, gcd(b, a % b))
    return gcd(b, a % b)


k = int(input())
sum = 0
dp = [[-1 for j in range(k+1)] for i in range(k+1)]
for i in range(1, k+1):
    for j in range(1, k+1):
        if dp[i][j] == -1:
            dp[i][j] = gcd(i, j)
        for k in range(1, k+1):
            ij = dp[i][j]
            if dp[ij][k] == -1:
                dp[ij][k] == gcd(ij, k)
                print("i,j,ij,k,val:", i, j, ij, k, dp[ij][k])
            sum += dp[ij][k]
            print("sum:", sum)
print(sum)

# def gcd(a, b: int) -> int:
#    if b == 0:
#        return a
#    return gcd(b, a % b)
#
#
# n = int(input())
# sum = 0
# dp = [[-1 for j in range(n)] for i in range(n)]
# for i in range(1, n+1):
#    for j in range(1, n+1):
#        dp[i-1][j-1] = gcd(i, j)
#        for k in range(1, n+1):
#            sum += gcd(dp[i-1][j-1], k)
# print(sum)
