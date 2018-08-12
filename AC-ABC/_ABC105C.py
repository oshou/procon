N = int(input())

max = 0
while True:
    if abs((-2) ** max) > N:
        break
    else:
        max += 1

arr = [ 0 for _ in range(max) ]

total = N
print(arr)
for i in range(0, max)[::-1]:
    print(i)
    if abs((-2) ** i) <= N:
        arr[i] = 1
        total -= (-2) ** i
        print("total: ", total)

for i in range(0, max)[--1]:
    print(arr[i])
