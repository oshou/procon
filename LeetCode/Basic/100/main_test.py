from main import Solution
from main import TreeNode


class TestSolution:
    def test_isSameTree(self):
        solution = Solution()
        cases = [
            [[1, 2, 3], [1, 2, 3], True],
            [[1, 2], [1, None, 2], False],
            [[1, 2, 1], [1, 1, 2], False],
        ]
        for p1, q1, expected in cases:
            p = TreeNode(p1[0], TreeNode(p1[1], None, None),
                         TreeNode(p1[2], None, None))
            q = TreeNode(q1[0], TreeNode(q1[1], None, None),
                         TreeNode(q1[2], None, None))
            assert solution.isSameTree(p, q) == expected
