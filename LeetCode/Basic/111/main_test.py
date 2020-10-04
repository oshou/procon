from main import TreeNode
from main import Solution


class TestSolution:
    def test_minDepth(self):
        solution = Solution()
        cases = [
            [[1, 2, 3, 4, 5], 2]
            #[[1, None, 2], 2]
            #[[3, 9, 20, None, None, 15, 7], 2]
        ]
        for args, expected in cases:
            if args[0]:
                root = TreeNode(args[0])
            if args[1]:
                root.left = TreeNode(args[1])
            if args[2]:
                root.right = TreeNode(args[2])
            if args[3]:
                root.left.left = TreeNode(args[3])
            if args[4]:
                root.left.right = TreeNode(args[4])
            # if args[5]:
            #    root.right.left = TreeNode(args[5])
            # if args[6]:
            #    root.right.right = TreeNode(args[6])

            assert solution.minDepth(root) == expected
