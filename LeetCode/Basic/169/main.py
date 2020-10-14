from typing import List
from collections import Counter


class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        m = Counter()
        for num in nums:
            m[num] += 1

        ans = 0
        size = len(nums)//2
        for k, v in m.items():
            if size < v:
                ans = k

        return ans
