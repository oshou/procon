a, b, c = map(int, input().split())
MOD = 998244353
print(a % MOD*(a+1) % MOD * b % MOD*(b+1) % MOD*c % MOD*(c+1) % MOD)
