from b import Solution


class TestSolution:
    def test_numPoints(self):
        solution = Solution()
        cases = [
            [4, 5, [0, -2, 3, 4], [5, 4, 4, -4], 3],
            [12, 3, [1, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3],
                [1, 1, 1, 1, 2, 3, 1, 2, 3, 1, 2, 3], 7],
        ]
        for n, d, x, y, expected in cases:
            assert solution.numPoints(n, d, x, y) == expected
