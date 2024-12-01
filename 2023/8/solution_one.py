import re

start = "AAA"
end = "ZZZ"
direction_to_idx = {"L": 0, "R": 1}

nodes = {}
directions = ""

with open("input.txt") as f:
    line_num = 0
    for line in f:
        if line_num == 0:
            directions = line.strip()
            line_num += 1
            continue
        if line_num == 1:
            line_num += 1
            continue
        node_re = re.search(r"(\w{3}) = \((\w{3}), (\w{3})\)", line)
        nodes[node_re.group(1)] = [node_re.group(2), node_re.group(3)]
        line_num += 1

print(directions)
print(nodes)

current = start
dir_idx = 0
steps = 0

while current != end:
    direction = direction_to_idx[directions[dir_idx]]
    current = nodes[current][direction]
    dir_idx += 1
    if dir_idx == len(directions):
        dir_idx = 0
    steps += 1

print(steps)
