N = int(input())
F = sum([int(x) for x in str(N)])
if N % F == 0:
    print("Yes")
else:
    print("No")
