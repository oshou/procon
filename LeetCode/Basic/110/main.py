class TreeNode:
    def __init__(self, val=0):
        self.val = val
        self.left = None
        self.right = None


class Solution:
    def isBalanced(self, root: TreeNode) -> bool:
        def maxDepth(root: TreeNode) -> int:
            if not root:
                return 0
            return 1+max(maxDepth(root.left), maxDepth(root.right))

        if not root:
            return False
        Abs = abs(maxDepth(root.left)-maxDepth(root.right))
        print("Abs:", Abs)
        if Abs > 1:
            return False
        return True
