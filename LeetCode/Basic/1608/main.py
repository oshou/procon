from typing import List


class Solution:
    def specialArray(self, nums: List[int]) -> int:
        max = -1
        for idx in range(0, len(nums)+1):
            cnt = 0
            for num in nums:
                if num >= idx:
                    cnt += 1
            if idx == cnt and max < cnt:
                max = cnt
        return max
