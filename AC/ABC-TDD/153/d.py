def sep(h: int) -> int:
    if h == 1:
        return 1
    return 1+2*sep(h//2)


h = int(input())
print(sep(h))
