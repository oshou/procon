from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        cnt = 0
        for i in range(0, len(nums)):
            if i == 0 or nums[i] != nums[i-1]:
                nums[cnt] = nums[i]
                cnt += 1
        return cnt
