from collections import deque

graph = []

with open('input.txt', 'r') as f:
    for line in f:
        graph.append([c for c in line.strip()])


start = (-1, -1)
for i in range(len(graph)):
    for j in range(len(graph[i])):
        if graph[i][j] == 'S':
            start = (i, j)
            break

print(start)

queue = deque([start])
steps = 64
while steps > 0:
    print("Step", steps)
    n = len(queue)
    for _ in range(n):
        i, j = queue.popleft()
        if i > 0 and graph[i-1][j] != '#' and (i-1, j):
            queue.append((i-1, j))
        if i < len(graph)-1 and graph[i+1][j] != '#':
            queue.append((i+1, j))
        if j > 0 and graph[i][j-1] != '#':
            queue.append((i, j-1))
        if j < len(graph[i])-1 and graph[i][j+1] != '#':
            queue.append((i, j+1))

        deduped_queue = list(dict.fromkeys(queue))
        queue = deque(deduped_queue)

    steps -= 1

places = set()
for place in queue:
    places.add(place)

print(len(places))

# for tile in places:
#     graph[tile[0]][tile[1]] = 'O'

# for row in graph:
#     print(''.join(row))
