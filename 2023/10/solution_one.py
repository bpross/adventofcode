from collections import defaultdict
from collections import deque

NORTH = 0
SOUTH = 1
EAST = 2
WEST = 3

lines = []
with open("input.txt", "r") as f:
    for line in f:
        lines.append([*line.strip()])

start = None
neighbors = defaultdict(list) 
distances = defaultdict(int)

for r in range(len(lines)):
    for c in range(len(lines[r])):
        if lines[r][c] == "S":
            start = (r, c)

# figure out which directions to go from start
# use BFS in those two directions
# once the pointers each each other, return the number of steps
r, c = start

def get_start_neighbors(r, c):
    north, south, east, west = (r - 1, c), (r + 1, c), (r, c + 1), (r, c - 1)
    north_pipe, south_pipe, east_pipe, west_pipe = "", "", "", ""
    
    if north[0] >= 0:
        north_pipe = lines[north[0]][north[1]]
    if south[0] < len(lines):
        south_pipe = lines[south[0]][south[1]]
    if east[1] < len(lines[r]):
        east_pipe = lines[east[0]][east[1]]
    if west[1] >= 0:
        west_pipe = lines[west[0]][west[1]]
    
    neighbors = []
    if north_pipe in ["7", "F", "|"]:
        if north_pipe == "7":
            neighbors.append((north[0], north[1], WEST))
        if north_pipe == "F":
            neighbors.append((north[0], north[1], EAST))
        if north_pipe == "|":
            neighbors.append((north[0], north[1], NORTH))
    if south_pipe in ["L", "J", "|"]:
        if south_pipe == "L":
            neighbors.append((south[0], south[1], EAST))
        if south_pipe == "J":
            neighbors.append((south[0], south[1], WEST))
        if south_pipe == "|":
            neighbors.append((south[0], south[1], SOUTH))
    if east_pipe in ["7", "J", "-"]:
        if east_pipe == "7":
            neighbors.append((east[0], east[1], SOUTH))
        if east_pipe == "J":
            neighbors.append((east[0], east[1], NORTH))
        if east_pipe == "-":
            neighbors.append((east[0], east[1], EAST))
    if west_pipe in ["-", "F", "L"]:
        if west_pipe == "-":
            neighbors.append((west[0], west[1], WEST))
        if west_pipe == "F":
            neighbors.append((west[0], west[1], SOUTH))
        if west_pipe == "L":
            neighbors.append((west[0], west[1], NORTH))

    return neighbors


def get_neighbors(r, c, direction):
    north, south, east, west = (r - 1, c), (r + 1, c), (r, c + 1), (r, c - 1)
    north_pipe, south_pipe, east_pipe, west_pipe = "", "", "", ""
    
    if north[0] >= 0:
        north_pipe = lines[north[0]][north[1]]
    if south[0] < len(lines):
        south_pipe = lines[south[0]][south[1]]
    if east[1] < len(lines[r]):
        east_pipe = lines[east[0]][east[1]]
    if west[1] >= 0:
        west_pipe = lines[west[0]][west[1]]

    neighbors = []
    if north_pipe in ["7", "F", "|"] and direction == NORTH:
        if north_pipe == "7":
            neighbors.append((north[0], north[1], WEST))
        elif north_pipe == "F":
            neighbors.append((north[0], north[1], EAST))
        elif north_pipe == "|":
            neighbors.append((north[0], north[1], NORTH))
    if south_pipe in ["L", "J", "|"] and direction == SOUTH:
        if south_pipe == "L":
            neighbors.append((south[0], south[1], EAST))
        elif south_pipe == "J":
            neighbors.append((south[0], south[1], WEST))
        elif south_pipe == "|":
            neighbors.append((south[0], south[1], SOUTH))
    if east_pipe in ["7", "J", "-"] and direction == EAST:
        if east_pipe == "7":
            neighbors.append((east[0], east[1], SOUTH))
        elif east_pipe == "J":
            neighbors.append((east[0], east[1], NORTH))
        elif east_pipe == "-":
            neighbors.append((east[0], east[1], EAST))
    if west_pipe in ["-", "F", "L"] and direction == WEST:
        if west_pipe == "-":
            neighbors.append((west[0], west[1], WEST))
        elif west_pipe == "F":
            neighbors.append((west[0], west[1], SOUTH))
        elif west_pipe == "L":
            neighbors.append((west[0], west[1], NORTH))
    
    return neighbors

neighbors = deque(get_start_neighbors(r, c))
farthest = 1
visited = set()
found = False

while len(neighbors) > 0:
    n = len(neighbors)
    for i in range(n):
        node = neighbors.popleft()
        visited.add(node)
        for neighbor in get_neighbors(node[0], node[1], node[2]):
            if neighbor in visited:
                found = True
                break
            neighbors.append(neighbor)
    farthest += 1
    if found:
        break

print(farthest // 2)
