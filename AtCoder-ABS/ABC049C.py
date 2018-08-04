S = str(input())
given = ['dream', 'dreamer', 'erase', 'eraser']

while len(S) > 0:
    judge = [S.endswith(i) for i in given]
    if any(judge):
        target = given[judge.index(True)]
        S = S[:-len(target)]
        if len(S) == 0:
            print("Yes")
            break
    else:
        print("No")
        break
