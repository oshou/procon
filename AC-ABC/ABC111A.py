N = str(input())
ans = ""
for n in N:
    if n == "1":
        ans += "9"
    elif n == "9":
        ans += "1"
    else:
        ans += n

print(ans)
