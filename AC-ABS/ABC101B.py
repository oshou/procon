N = int(input())
S = sum([int(x) for x in str(N)])
if N % S == 0:
    print("Yes")
else:
    print("No")
