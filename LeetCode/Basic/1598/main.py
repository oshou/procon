from typing import List


class Solution:
    def minOperations(self, logs: List[str]) -> int:
        cnt = 0
        for log in logs:
            if log == "../":
                if cnt != 0:
                    cnt -= 1
            elif log == "./":
                cnt -= 0
            else:
                cnt += 1

        return cnt
