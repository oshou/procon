class UnionFind():
    def __init__(self, n):
        self.n = n
        self.parents = [-1]*n
        for i in range(n):
            self.parents[i] = i

    # 要素xが所属するグループの根を返す
    def root(self, x) -> int:
        if self.parents[x] == x:
            return x
        else:
            self.parents[x] = self.root(self.parents[x])
            return self.parents[x]

    def unite(self, x, y):
        x = self.root(x)
        y = self.root(y)
        if x == y:
            return
        elif x > y:
            u.parents[x] = y
        elif x < y:
            u.parents[y] = x

    def same(self, x, y) -> bool:
        return self.root(x) == self.root(y)


n, m = map(int, input().split())
cnt = n
u = UnionFind(n)
for i in range(m):
    a, b = map(int, input().split())
    if u.same(a-1, b-1):
        continue
    else:
        u.unite(a-1, b-1)
        cnt -= 1
if cnt > 0:
    print(cnt-1)
else:
    print(cnt)
