from main import Solution


class TestSolution:
    def test_reverse(self):
        solution = Solution()
        cases = [
            [123, 321],
            [-123, -321],
            [120, 21],
            [1534236469, 0]
        ]
        for arg, expected in cases:
            assert solution.reverse(arg) == expected
