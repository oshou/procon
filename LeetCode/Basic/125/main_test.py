from main import Solution


class TestSolution:
    def test_isPalindrome(self):
        solution = Solution()
        cases = [
            ["A man, a plan, a canal: Panama", True],
            ["race a car", False],
            ["0P", False]
        ]
        for s, expected in cases:
            assert solution.isPalindrome(s) == expected
