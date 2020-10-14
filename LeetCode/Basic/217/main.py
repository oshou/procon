from typing import List
from collections import Counter


class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        m = Counter()
        for num in nums:
            if m[num]:
                return True
            else:
                m[num] = True
        return False
