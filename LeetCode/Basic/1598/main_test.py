from main import Solution


class TestSolution:
    def test_minOperations(self):
        solution = Solution()
        cases = [
            [["d1/", "d2/", "../", "d21/", "./"], 2],
            [["d1/", "d2/", "./", "d3/", "../", "d31/"], 3],
            [["d1/", "../", "../", "../"], 0],
            [["./", "../", "./"], 0]
        ]
        for logs, expected in cases:
            assert solution.minOperations(logs) == expected
