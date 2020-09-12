from main import Solution


class TestSolution:
    def test_romanToInt(self):
        solution = Solution()
        cases = [
            ["III", 3],
            ["IV", 4],
            ["IX", 9],
            ["LVIII", 58],
            ["MCMXCIV", 1994],
        ]
        for arg, expected in cases:
            assert solution.romanToInt(arg) == expected
