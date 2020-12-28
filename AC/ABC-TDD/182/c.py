def getDiv3Cnt(s: str):
    if len(s) == 0:
        return -1
    if int(s) % 3 == 0:
        return 0

    mincnt = len(s)+1
    for i in range(len(s)):
        div = getDiv3Cnt(s[:i]+s[i+1:])
        if div != -1:
            mincnt = min(mincnt, 1+div)
    if mincnt == len(s)+1:
        return -1
    else:
        return mincnt
