from main import Solution


class TestSolution:
    def test_majorityElement(self):
        solution = Solution()
        cases = [
            [[3, 2, 3], 3],
            [[2, 2, 1, 1, 1, 2, 2], 2]
        ]
        for nums, expected in cases:
            assert solution.majorityElement(nums) == expected
