from main import Solution


class TestSolution:
    def test_isPalindrome(self):
        solution = Solution()
        cases = [
            [121, True],
            [-121, False],
            [10, False]
        ]
        for arg, expected in cases:
            assert solution.isPalindrome(arg) == expected
