def canFetch(s: str) -> bool:
    if len(s) == 0:
        return True
    if len(s) >= 5 and (s.startswith("dream") or s.startswith("erase")):
        if canFetch(s[5:]):
            return True
    if len(s) >= 6 and s.startswith("eraser"):
        if canFetch(s[6:]):
            return True
    if len(s) >= 7 and s.startswith("dreamer"):
        if canFetch(s[7:]):
            return True
    return False


s = input()
if canFetch(s):
    print("YES")
else:
    print("NO")
