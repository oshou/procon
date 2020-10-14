from main import Solution


class TestSolution:
    def test_maxSubArray(self):
        solution = Solution()
        cases = [
            [[-2, 1, -3, 4, -1, 2, 1, -5, 4], 6],
            [[1], 1],
            [[0], 0],
            [[-1], -1],
            [[-2147483647], -2147483647],
            [[-2, -1], -1]
        ]
        for nums, expected in cases:
            assert solution.maxSubArray(nums) == expected
