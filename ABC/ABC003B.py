S = input()
T = input()
candidate = ["@","a","t","c","o","d","e","r"]

judge = "You can win"
for i in range(len(S)):
    if S[i] == T[i]:
        pass
    else:
        if S[i] in candidate and T[i] in candidate:
            pass
        else:
            judge = "You will lose"

print(judge)
