N = int(input())
seats = 0
for i in range(N):
    l, r = map(int, input().split())
    seats += r - l + 1

print(seats)
