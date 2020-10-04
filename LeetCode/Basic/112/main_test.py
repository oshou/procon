from pprint import pprint
from main import TreeNode
from main import Solution


class TestSolution:
    def test_hasPathSum(self):
        solution = Solution()
        cases = [
            [[5, 4, 8, 11, 13, 4, 7, 2, 1], 22, True]
        ]
        for args, sum, expected in cases:
            root = TreeNode(args[0])
            root.left = TreeNode(args[1])
            root.right = TreeNode(args[2])
            root.left.left = TreeNode(args[3])
            root.right.left = TreeNode(args[4])
            root.right.right = TreeNode(args[5])
            root.left.left.left = TreeNode(args[6])
            root.left.left.right = TreeNode(args[7])
            root.right.right.right = TreeNode(args[8])

            assert solution.hasPathSum(root, sum) == expected
