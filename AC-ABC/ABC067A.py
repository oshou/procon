A, B = map(int, input().split())

if not (A % 3) and (B % 3) and ((A+B) % 3):
    print("Possible")
else:
    print("Impossible")
