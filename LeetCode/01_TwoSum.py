class Solution:
    def twoSum(self, nums, target):
        """
        BruteForce
        - TimeComplexity: O(n^2)
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for i in range(len(nums)):
            mid = target-nums[i]
            for j in range(i+1, len(nums)):
                if nums[j] == mid:
                    return [i, j]


if __name__ == "__main__":
    nums = [1, 4, 8, 3, 2, 9, 15]
    target = 13
    sol = Solution()
    print("solution: ", sol.twoSum(nums, target))
