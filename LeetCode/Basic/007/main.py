class Solution:
    def reverse(self, x: int) -> int:
        num, reversed = 0, 0
        is_negative = False

        if x < 0:
            is_negative = True

        num = abs(x)
        while num != 0:
            reversed = reversed*10+num % 10
            num /= 10

        if is_negative:
            return (-1)*reversed

        return reversed
