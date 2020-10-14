A = int(input())
S = str(input())

ans = "No"
val = A

if A == 0:
    ans = "Yes"
    break
else:
    for i in S:
        if i == "+":
            val += 1
        elif i == "-":
            val -= 1
        else:
            pass

        if val == 0:
            ans = "Yes"
            break

print(ans)
