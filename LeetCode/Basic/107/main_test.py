from main import TreeNode
from main import Solution


class TestSolution:
    def test_levelOrderBottom(self):
        solution = Solution()
        cases = [
            [[3, 9, 20, None, None, 15, 7], 3]
        ]
        for args, expected in cases:
            if args[0] is not None:
                root = TreeNode(args[0])
            if args[1] is not None:
                root.left = TreeNode(args[1])
            if args[2] is not None:
                root.right = TreeNode(args[2])
            if args[3] is not None:
                root.left.left = TreeNode(args[3])
            if args[4] is not None:
                root.left.right = TreeNode(args[4])
            if args[5] is not None:
                root.right.left = TreeNode(args[5])
            if args[6] is not None:
                root.right.right = TreeNode(args[6])

            assert solution.levelOrderBottom(root)
