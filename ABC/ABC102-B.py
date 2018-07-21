def compose():
    return lambda x: abs(int(x))


n = int(input())
if 2 <= n <= 100:
    A = list(map(compose(), input().split()))
    print(max(A)-min(A))
