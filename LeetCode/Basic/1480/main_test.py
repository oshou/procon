from main import Solution


class TestSolution:
    def test_runningSum(self):
        solution = Solution()
        cases = [
            [[1, 2, 3, 4], [1, 3, 6, 10]],
            [[1, 1, 1, 1, 1], [1, 2, 3, 4, 5]],
            [[3, 1, 2, 10, 1], [3, 4, 6, 16, 17]],
        ]
        for nums, expected in cases:
            assert solution.runningSum(nums) == expected
