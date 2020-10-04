from main import ListNode
from main import Solution


class TestSolution:
    def __eq__(self, other):
        if isinstance(self, other.__class__):

    def test_addTwoNumbers(self):
        solution = Solution()
        l1 = ListNode(2)
        l1.next = ListNode(4)
        l1.next.next = ListNode(3)
        l2 = ListNode(5)
        l2.next = ListNode(6)
        l2.next.next = ListNode(4)
        lsum = ListNode(7)
        lsum.next = ListNode(0)
        lsum.next.next = ListNode(8)
        solution.addTwoNumbers(l1, l2) == lsum
