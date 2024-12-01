from collections import deque

UP = "UP"
DOWN = "DOWN"
RIGHT = "RIGHT"
LEFT = "LEFT"

graph = []

with open("sample.txt", "r") as f:
    for line in f:
        graph.append([int(c) for c in line.strip()])


def reached_end(path):
    if path[0]:
        return True
    return False

def dfs(r, c, visited):
    if r < 0 or r >= len(graph) or c < 0 or c >= len(graph[0]):
        return False, -1
    if r == len(graph) - 1 and c == len(graph[0]) - 1:
        # Reached the end
        return True, 0 
    if (r,c) in visited:
        return False, - 1

    visited.add((r, c))

    up_result = dfs(r - 1, c, visited)
    down_result = dfs(r + 1, c, visited)
    right_result = dfs(r, c + 1, visited)
    left_result = dfs(r, c - 1, visited)
    paths = [up_result, down_result, right_result, left_result]

    paths = list(filter(reached_end, paths))
    if len(paths) == 0:
        # none reached the end
        return False, - 1

    val = min(paths, key = lambda t: t[1])
    visited.remove((r, c))
    return True, val[1] + graph[r][c] 

visited = set()
right, right_heat_loss = dfs(0, 0, visited)
visited = set()
down, down_heat_loss = dfs(0, 0, visited)

print(right, right_heat_loss, down, down_heat_loss)
