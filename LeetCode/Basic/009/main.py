class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        strx = str(x)
        if strx == strx[::-1]:
            return True

        return False
