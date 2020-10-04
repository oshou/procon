from main import Solution


class TestSolution:
    def test_kidsWithCandies(self):
        solution = Solution()
        cases = [
            [[2, 3, 5, 1, 3], 3, [True, True, True, False, True]],
            [[4, 2, 1, 1, 2], 1, [True, False, False, False, False]],
            [[12, 1, 12], 10, [True, False, True]],
            [[2, 8, 7], 1, [False, True, True]],
        ]

        for candies, extraCandies, expected in cases:
            assert solution.kidsWithCandies(candies, extraCandies) == expected
