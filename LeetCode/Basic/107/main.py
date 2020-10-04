from typing import List


class TreeNode:
    def __init__(self, val: int = 0):
        self.val = val
        self.left = None
        self.right = None


class Solution:
    def levelOrderBottom(self, root: TreeNode) -> List[List[int]]:
        return None

    def inorder(node: TreeNode, depth: int, traversal: List[int]):
        if not node:
            return
        if len(traversal) == depth:
            traversal.append([])
        self.inorder(node.left, depth+1, traversal)
        traversal.append(node.val)
        self.inorder(node.right, depth+1, traversal)
