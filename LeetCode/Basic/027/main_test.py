from main import Solution


class TestSolution:
    def test_removeElement(self):
        solution = Solution()
        cases = [
            [[3, 2, 2, 3], 3, 2],
            [[0, 1, 2, 2, 3, 0, 4, 2], 2, 5],
        ]
        for nums, val, expected in cases:
            assert solution.removeElement(nums, val) == expected
