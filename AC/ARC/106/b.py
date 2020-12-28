class Graph:
    def __init__(self):
        self.dict = {}

    def add_vertex(self, v):
        self.dict[v] = []

    def add_edge(self, v1, v2):
        self.dict[v1].append(v2)
        self.dict[v2].append(v1)

    def graph(self):
        return self.dict

    def entry(self, v1):
        return self.dict[v1]


n, m = map(int, input().split())
a = list(map(int, input().split()))
b = list(map(int, input().split()))

diff = [b[i]-a[i] for i in range(n)]


graph = Graph()
for i in range(n):
    graph.add_vertex(i)

for i in range(m):
    tmpc, tmpd = map(int, input().split())
    graph.add_edge(tmpc-1, tmpd-1)

for i in range(n):
    if diff[i] == (-1)*sum(diff[i] for i in graph.entry(i)):
        print("Yes")
        exit()
print("No")
