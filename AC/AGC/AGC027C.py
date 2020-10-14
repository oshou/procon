N, M = map(int, input().split())
s = input()
a = []
b = []
A_path = []
B_path = []
for i in range(M):
    tmpa, tmpb = map(int, input().split())
    a.append(tmpa)
    b.append(tmpb)

for i in range(M):
    if s[a[i]-1] == s[b[i]-1]:
        if s[a[i]-1] == "A":
            A_path.append([a[i], b[i]])
        elif s[a[i]-1] == "B":
            B_path.append([a[i], b[i]])

for ele1 in A_path:
    for ele2 B_path:
        for i in range(M):
        if (ele1[1] == ele2[1] and ele1[2] = ele2[2]) or




if A_path == 1 and B_path == 1:
    print("Yes")
else:
    print("No")
