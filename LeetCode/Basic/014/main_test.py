from main import Solution


class TestSolution:
    def test_longestCommonPrefix(self):
        solution = Solution()
        cases = [
            [["flower", "flow", "flight"], "fl"],
            [["dog", "racecar", "car"], ""],
            [[], ""],
            [[""], ""],
            [["a"], "a"]
        ]
        for args, expected in cases:
            assert solution.longestCommonPrefix(args) == expected
