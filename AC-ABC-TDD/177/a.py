def is_intime(d, t, s: int) -> bool:
    if s*t >= d:
        return True
    else:
        return False


d, t, s = map(int, input().split())
if is_intime(d, t, s):
    print("Yes")
else:
    print("No")
