from main import Solution


class TestSolution:
    def test_removeDuplicates(self):
        solution = Solution()
        cases = [
            [[1, 1, 2], 2],
            [[0, 0, 1, 1, 1, 2, 2, 3, 3, 4], 5]
        ]
        for nums, expected in cases:
            assert solution.removeDuplicates(nums) == expected
