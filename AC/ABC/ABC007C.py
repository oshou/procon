R, C = map(int, input().split())
sy, sx = map(int, input().split())
gy, gx = map(int, input().split())
c = [list(input()) for _ in range(R)]
steps = [[-1]*C for _ in range(R)]
dxy = ((1, 0), (0, 1), (-1, 0), (0, -1))

def bfs(startx, starty):
    queue = []
    queue.insert(0, (startx, starty))
    steps[starty-1][startx-1] = 0

    while(len(queue)):
        x, y = queue.pop()
        for (dx, dy) in dxy:
            nx, ny = x + dx, x + dy

            if (c[ny-1][nx-1] != "#" and steps[ny-1][nx-1] < 0):
                queue.insert(0, (nx, ny))
                steps[ny-1][nx-1] = steps[y-1][x-1] + 1
                print("steps", ny-1, nx-1, ": ", steps[ny-1][nx-1])

bfs(sx, sy)
print(steps[gy-1][gx-1])
