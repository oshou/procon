s = input()
s1i = int(s[0:2])
s2i = int(s[2:4])
isYYMM = False
isMMYY = False
if 1 <= s1i and s1i <= 12:
    isMMYY = True
if 1 <= s2i and s2i <= 12:
    isYYMM = True

if isYYMM and isMMYY:
    print("AMBIGUOUS")
elif isYYMM:
    print("YYMM")
elif isMMYY:
    print("MMYY")
else:
    print("NA")
