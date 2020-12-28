def min_maruta(n:int)->int:
    if n <= 2:
        return 1
    elif n == 3:
        return 2
    else:
        return 2 + (n-3)

n = int(input())
print(min_maruta(n))
