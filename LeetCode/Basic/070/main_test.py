from main import Solution


class TestSolution:
    def test_climbStairs(self):
        solution = Solution()
        cases = [
            [2, 2],
            [3, 3],
            [5, 8],
        ]
        for n, expected in cases:
            assert solution.climbStairs(n) == expected
