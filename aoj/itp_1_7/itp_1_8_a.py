line = input()
ans = ""
for i in range(len(line)):
    if (line[i].islower()):
        ans += str(line[i].upper)
    elif (line[i].isupper):
        ans += str(line[i].lower)

print(ans)
