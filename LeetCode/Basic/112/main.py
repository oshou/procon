class TreeNode:
    def __init__(self, val=0):
        self.val = val
        self.left = None
        self.right = None


class Solution:
    def hasPathSum(self, root: TreeNode, sum: int) -> bool:
        if not root:
            return False
        sum -= root.val
        if sum == 0 and not root.left and not root.right:
            return True

        return self.hasPathSum(root.left, sum) or self.hasPathSum(root.right, sum)
