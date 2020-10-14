from main import Solution


class TestSolution:
    def test_specialArray(self):
        solution = Solution()
        cases = [
            [[3, 5], 2],
            [[0, 0], -1],
            [[0, 4, 3, 0, 4], 3],
            [[3, 6, 7, 7, 0], -1],
        ]
        for nums, expected in cases:
            assert solution.specialArray(nums) == expected
