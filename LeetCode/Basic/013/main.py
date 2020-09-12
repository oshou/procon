class Solution:
    def romanToInt(self, s: str) -> int:
        doubles = {
            "IV": 4,
            "IX": 9,
            "XL": 40,
            "XC": 90,
            "CD": 400,
            "CM": 900,
        }
        singles = {
            "I": 1,
            "V": 5,
            "X": 10,
            "L": 50,
            "C": 100,
            "D": 500,
            "M": 1000,
        }

        i, sum = 0, 0
        while i <= len(s)-1:
            if i < len(s)-1 and s[i:i+2] in doubles:
                sum += doubles[s[i:i+2]]
                i += 2
            else:
                sum += singles[s[i]]
                i += 1

        return sum
