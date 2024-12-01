from collections import deque

RIGHT = "RIGHT"
LEFT = "LEFT"
UP = "UP"
DOWN = "DOWN"

graph = []

with open("input.txt", "r") as f:
    for line in f:
        graph.append([c for c in line.strip()])


def get_neighbor(r, c, direction):
    if r < 0 or c < 0 or r >= len(graph) or c >= len(graph[0]):
        return
    char = graph[r][c]
    if char == ".":
        if direction == RIGHT:
            return [(r, c + 1, RIGHT)]
        elif direction == LEFT:
            return [(r, c - 1, LEFT)]
        elif direction == UP:
            return [(r - 1, c, UP)]
        else:
            return [(r + 1, c, DOWN)]
    elif char == "|":
        if direction == RIGHT or direction == LEFT:
            return [(r - 1, c, UP), (r + 1, c, DOWN)]
        elif direction == UP:
            return [(r - 1, c, UP)]
        else:
            return [(r + 1, c, DOWN)]
    elif char == "-":
        if direction == RIGHT:
            return [(r, c + 1, RIGHT)]
        elif direction == LEFT:
            return [(r, c - 1, LEFT)]
        elif direction == UP or direction == DOWN:
            return [(r, c - 1, LEFT), (r, c + 1, RIGHT)]
    elif char == "/":
        if direction == RIGHT:
            return [(r - 1, c, UP)]
        elif direction == LEFT:
            return [(r + 1, c, DOWN)]
        elif direction == UP:
            return [(r, c + 1, RIGHT)]
        elif direction == DOWN:
            return [(r, c - 1, LEFT)]
    elif char == "\\":
        if direction == RIGHT:
            return [(r + 1, c, DOWN)]
        elif direction == LEFT:
            return [(r - 1, c, UP)]
        elif direction == UP:
            return [(r, c - 1, LEFT)]
        elif direction == DOWN:
            return [(r, c + 1, RIGHT)]

    return []


visited = set()
neighbors = deque([(0, 0, RIGHT)])
visited.add((0,0,RIGHT))
while neighbors:
    n = len(neighbors)
    for _ in range(n):
        node = neighbors.popleft()
        for neighbor in get_neighbor(*node):
            r, c, direction = neighbor
            if r < 0 or r >= len(graph) or c < 0 or c >= len(graph):
                continue
            if neighbor in visited:
                continue
            visited.add(neighbor)
            neighbors.append(neighbor)

energized = set()
for node in visited:
    energized.add((node[0], node[1]))

print(len(energized))
