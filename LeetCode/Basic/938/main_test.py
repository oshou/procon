from main import TreeNode
from main import Solution


class TestSolution:
    def test_rangeSumBST(self):
        solution = Solution()
        cases = [
            [[10, 5, 15, 3, 7, None, 18], 7, 15, 32],
            [[10, 5, 15, 3, 7, 13, 18, 1, None, 6], 6, 10, 23]
        ]
        for root, l, r, expected in cases:
            if root[0]:
                root = TreeNode(root[0])
            if root[1]:
                root.left = TreeNode(root[1])
            if root[2]:
                root.right = TreeNode(root[2])
            if root[3]:
                root.left.left = TreeNode(root[3])
            if root[4]:
                root.left.right = TreeNode(root[4])
            if root[5]:
                root.right.left = TreeNode(root[5])
            if root[6]:
                root.right.right = TreeNode(root[6])
