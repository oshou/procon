from typing import List


class NodeTree:
    def __init__(self, val=0):
        self.val = val
        self.left = None
        self.right = None


class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> TreeNode:
        root = TreeNode()
