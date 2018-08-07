import math

m = int(input())
if m < 100:
    print("00")
elif m < 1000:
    print("0", round(m/100), sep="")
elif m <= 5000:
    print(round(m/100))
elif m <= 30000:
    print(round(m/1000)+50)
elif m <= 70000:
    print(round((m/1000-30)/5) + 80)
elif 70000 < m:
    print(89)

