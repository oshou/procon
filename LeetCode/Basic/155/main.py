class MinStack:
    def __init__(self):
        self.main = []
        self.mins = []

    def push(self, x: int) -> None:
        self.main.append(x)
        if not self.mins or x <= self.mins[-1]:
            self.mins.append(x)

    def pop(self) -> None:
        item = self.main.pop()
        if item == self.mins[-1]:
            self.mins.pop()

    def top(self) -> int:
        return self.main[-1]

    def getMin(self) -> int:
        return self.mins[-1]
