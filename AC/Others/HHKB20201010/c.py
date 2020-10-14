from typing import List


def min(arr: List[int], hit: int) -> int:
    size = len(arr)
    for j in range(0, size):
        if arr[j] >= 0:
            return j


n = int(input())
p = list(map(int, input().split()))
MAX = 2000001
arr = [1]*MAX
mmin = 0
for i in range(mmin, n):
    arr[p[i]] = -1
    if arr[mmin] < 0:
        mmin = min(arr, p[i])
    print(mmin)
