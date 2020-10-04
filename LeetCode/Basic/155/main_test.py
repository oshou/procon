from main import MinStack


class TestSolution:
    def test_push(self):
        cases = [
            ["push", -2, None],
            ["push", 0, None],
            ["push", -3, None],
            ["getMin", None, -3],
            ["pop", None, None],
            ["top", None, 0],
            ["getMin", None, -2]
        ]
        minStack = MinStack()
        for method, arg, expected in cases:
            minStack.method(arg) == expected
