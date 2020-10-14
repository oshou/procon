from main import Solution


class TestSolution:
    def test_calc(self):
        solution = Solution()
        cases = [
            [2, 14],
            [10, 1110]
        ]
        for arg, expected in cases:
            assert solution.calc(arg) == expected
