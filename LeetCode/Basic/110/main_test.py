from main import TreeNode
from main import Solution


class TestSolution:
    def test_isBalanced(self):
        solution = Solution()
        cases = [
            [[3, 9, 20, None, None, 15, 7], True],
            [[1, 2, 2, 3, 3, None, None, 4, 4], False]
        ]
        for nums, expected in cases:
            solution.insert(root, nums)
            print(root.val)
            print(root.left.val)
            print(root.right.val)
            print(root.left.left.val)
            print(root.left.right.val)
            print(root.right.left.val)
            print(root.right.right.val)
            solution.insert(root, nums)
            assert solution.isBalanced(root) == expected
