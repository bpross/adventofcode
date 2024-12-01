import re

lines = []
with open("input.txt", "r") as f:
    for line in f:
        lines.append(line)


time = ''
distance = ''

time_groups = re.search(r"Time: (.*)", lines[0])
distance_groups = re.search(r"Distance: (.*)", lines[1])

for t in time_groups.group(1).strip().split(" "):
    if t != '':
        time += t

time = int(time)

for d in distance_groups.group(1).strip().split(" "):
    if d != '':
        distance += d

distance = int(distance)

times = [time]
distances = [distance]

ways_to_win = []
for i in range(len(times)):
    starting_speed = 0
    ways_to_set_record = 0

    for j in range(times[i]):
        total_distance = starting_speed * (times[i] - j)
        if total_distance > distances[i]:
            ways_to_set_record += 1
        starting_speed += 1
    ways_to_win.append(ways_to_set_record)

total = -1
for w in ways_to_win:
    if w > 0:
        if total == -1:
            total = w
        else:
            total *= w

if total == -1:
    print(0)

print(total)
