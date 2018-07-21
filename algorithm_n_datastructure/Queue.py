class Queue:
    def __init__(self):
        head = tail = 0

    def isEmpty(self):
        return head == tail

    def isFull(self):
        return head == (tail + 1) % MAX

    def enqueue(x):
