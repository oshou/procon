from main import Solution


class TestSolution:
    def test_containsDuplicate(self):
        solution = Solution()
        cases = [
            [[1, 2, 3, 1], True],
            [[1, 2, 3, 4], False],
            [[1, 1, 1, 3, 3, 4, 3, 2, 4, 2], True]
        ]
        for nums, expected in cases:
            assert solution.containsDuplicate(nums) == expected
