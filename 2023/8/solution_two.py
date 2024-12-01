from copy import deepcopy
import math
import re

starts = []
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
        if node_re.group(1)[-1] == "A":
            starts.append(node_re.group(1))
        line_num += 1

print(starts)

node_steps = []

for start in starts:
    dir_idx = 0
    steps = 0
    end = False
    current = start
    while not end:
        direction = direction_to_idx[directions[dir_idx]]
        current = nodes[current][direction]
        dir_idx += 1
        if dir_idx == len(directions):
            dir_idx = 0
        steps += 1
        if current[-1] == "Z":
            end = True

    node_steps.append(steps)

print(node_steps)
print(math.lcm(*node_steps))
