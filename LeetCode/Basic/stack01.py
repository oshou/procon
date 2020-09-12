class Algo:
    def __init__(self):
        self.stack = []
        self.queue = []

    def push(self, x):
        self.stack.append(x)

    def pop(self):
        self.stack.pop(-1)

    def enqueue(self, x):
        self.queue.append(x)

    def dequeue(self):
        self.queue.pop(0)


algo = Algo()
x = int(input())

for i in range(x):
    algo.push(i)
    algo.enqueue(i)

print(algo.stack)
print(algo.queue)

for i in range(x):
    algo.pop()
    algo.dequeue()
    print('stack.pop', end="")
    print(algo.stack)
    print('queue.dequeue', end="")
    print(algo.queue)
    print("*********************************")
    print("*********************************")
