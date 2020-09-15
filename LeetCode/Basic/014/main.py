from typing import List


class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        minlen = 0
        if len(strs) == 0:
            return ""
        elif len(strs) == 1:
            return strs[0]

        for str in strs:
            if minlen == 0 or len(str) < minlen:
                minlen = len(str)
                minstr = str

        if len(minstr) == 0:
            return ""

        for i in range(len(minstr)):
            for str in strs:
                if not str.startswith(minstr[:i]):
                    return minstr[:i-1]
