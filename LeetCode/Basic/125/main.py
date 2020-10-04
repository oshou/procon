class Solution:
    def isPalindrome(self, s: str) -> bool:
        flatted = ""
        for c in s:
            if c.isalnum():
                flatted += c.lower()

        if flatted == flatted[::-1]:
            return True

        return False
