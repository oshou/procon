'''
/**
n: entry

b: buildings(max: 4)
f: floors(max: 3)
r: rooms(max: 10)
v: person
'''


# Initialize
house = [[[0 for r in range(10)] for f in range(3)] for b in range(4)]

# Input
n = int(input())
for i in range(n):
    b, f, r, v = map(int, input().split())
    house[b-1][f-1][r-1] += v

# Output
for i in range(4):
    for j in range(3):
        print(' '+' '.join(map(str, house[i][j])))

    if i != 3:
        print("#"*20)
