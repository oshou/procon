a, b = map(int, input().split())
n = b - a - 1
x = int(n*(n+1)/2 - a)
print(x)
