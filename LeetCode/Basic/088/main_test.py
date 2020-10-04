from main import Solution


class TestSolution:
    def test_merge(self):
        solution = Solution()
        cases = [
            [[1, 2, 3, 0, 0, 0], 3, [2, 5, 6], 3, [1, 2, 2, 3, 5, 6]]
        ]
        for nums1, m, nums2, n, expected in cases:
            solution.merge(nums1, m, nums2, n)
            assert nums1 == expected
