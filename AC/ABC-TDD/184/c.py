def ryuma(a, b, c, d) -> int:
    if a == c and b == d:
        return 0
    if (a+b == c+d) or (a-b == c-d) or (abs(a-c)+abs(b-d) <= 3):
        return 1
    else:
        print("X:", (b-a+c+d)/2)
        print("Y:", (c+d-b+a)/2)
        return 1+ryuma((b-a+c+d)/2, (c+d-b+a)/2, c, d)


a, b = map(int, input().split())
c, d = map(int, input().split())
print(ryuma(a, b, c, d))
