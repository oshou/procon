from main import Solution


class TestSolution:
    def test_twoSum(self):
        solution = Solution()
        cases = [
            [[2, 7, 11, 15], 9, [0, 1]],
            [[3, 2, 4], 6, [1, 2]],
            #[[3, 3], 6, [0, 1]],
        ]
        for arg, target, expected in cases:
            assert solution.twoSum(arg, target) == expected
