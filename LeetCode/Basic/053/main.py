from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        if not nums:
            return 0

        sums = []
        sum = 0
        for num in nums:
            sum += num
            sums.append(sum)
        nmax = max(sums)
        nmin = min(sums)
        print(sums)
        if nmax == nmin:
            return nmax
        else:
            return nmax-nmin
