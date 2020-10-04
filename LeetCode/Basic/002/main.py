class ListNode:
    def __init__(self, val=0):
        self.val = val
        self.next = None


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        prev = result = ListNode(None)
        carry = 0
        while l1 or l2 or carry:
            if l1:
                carry += l1.val
                l1 = l1.next
            if l2:
                carry += l2.val
                l2 = l2.next
            prev.next = ListNode(carry % 10)
            prev = prev.next
            carry //= 10

        return result.next
