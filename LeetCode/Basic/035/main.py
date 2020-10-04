from typing import List


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        for i in range(0, len(nums)):
            if target > nums[i]:
                if i == len(nums)-1:
                    return i+1
                continue
            else:
                return i
