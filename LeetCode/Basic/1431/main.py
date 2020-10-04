from typing import List


class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        max = 0
        for candy in candies:
            if max < candy:
                max = candy

        ans = [False]*len(candies)
        for i, candy in enumerate(candies):
            if candy + extraCandies >= max:
                ans[i] = True
            else:
                ans[i] = False
        return ans
