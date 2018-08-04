S = str(input())
T = str(input())

check = 'No'
for _ in range(len(S)):
    S = S[-1]+S[1::-1]
    if S == T:
        check = 'Yes'

print(check)
