class TreeNode:
    def __init__(self, val=0):
        self.val = val
        self.left = None
        self.right = None


class Solution:
    def minDepth(self, root: TreeNode) -> int:
        if not root:
            return 0
        if not root.left and not root.right:
            return 1
        minLeft = self.minDepth(root.left)
        minRight = self.minDepth(root.right)
        if not root.left:
            return 1 + minRight
        if not root.right:
            return 1 + minLeft

        return 1 + min(minLeft, minRight)
