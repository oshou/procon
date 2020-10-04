from typing import List


class Solution:
    def numPoints(self, n, d: int, x, y: List[int]) -> int:
        cnt = 0
        d2 = d*d
        for i in range(0, n):
            if d2 >= x[i]*x[i]+y[i]*y[i]:
                cnt += 1
        return cnt
