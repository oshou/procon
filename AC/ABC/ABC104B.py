import re

S = str(input())

ans = "AC"

uppercount = len(re.findall('[A-Z]', S))
if not uppercount == 2:
    ans = "WA"

if not S[0] == "A":
    ans = "WA"

cfind = S[2:-1].find("C")
if cfind < 0:
    ans = "WA"

if not S[1:cfind+2].islower and S[cfind+3:].islower():
    ans = "WA"

print(ans)
