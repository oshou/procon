class Solution:
    def is_valid(self, s: str) -> bool:
        match = {"}": "{", "]": "[", ")": "("}
        stack = []

        def is_empty(x):
            return x == []

        for c in s:
            print(stack)
            if c in ["[", "{", "("]:
                stack.append(c)
            else:
                if is_empty(stack):
                    return False
                if stack[-1] == match[c]:
                    stack.pop()
                else:
                    return False
        return is_empty(stack)


solution = Solution()
print(solution.is_valid("{}{"))
print(solution.is_valid("{[()]}"))
