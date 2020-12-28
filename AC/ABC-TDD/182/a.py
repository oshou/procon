a, b = map(int, input().split())
max = 2*a+100
if max <= b:
    print(0)
else:
    print(max-b)
