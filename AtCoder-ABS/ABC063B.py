S = str(input())
if sorted(list(S)) == sorted(set(S)):
    print("yes")
else:
    print("no")
