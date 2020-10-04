from main import Solution


class TestSolution:
    def test_twoSum(self):
        solution = Solution()
        cases = [
            [[2, 7, 11, 15], 9, [1, 2]],
            [[2, 3, 4], 6, [1, 3]],
            [[-1, 0], -1, [1, 2]]
        ]
        for nums, target, expected in cases:
            assert solution.twoSum(nums, target) == expected
