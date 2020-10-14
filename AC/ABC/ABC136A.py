def solve(a,b,c):
    return max(0,c-a+b)

a,b,c=map(int,input().split())
print(solve(a,b,c))