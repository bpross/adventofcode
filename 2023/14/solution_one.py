ROUND_ROCK = "O"
SQUARE_ROCK = "#"
OPEN = "."

platforms = []

with open('input.txt') as f:
    for line in f:
        platforms.append(line.strip())


def roll(top, bottom):
    new_top, new_bottom = [], []
    for i in range(len(top)):
        if top[i] == ROUND_ROCK or top[i] == SQUARE_ROCK or (top[i] == OPEN and bottom[i] == OPEN) or (top[i] == OPEN and bottom[i] == SQUARE_ROCK):
            new_top.append(top[i])
            new_bottom.append(bottom[i])
        elif top[i] == OPEN and bottom[i] == ROUND_ROCK:
            new_top.append(bottom[i])
            new_bottom.append(OPEN)
        else:
            print(top[i], bottom[i])

    return new_top, new_bottom


def check_platforms(top, bottom):
    for i in range(len(top)):
        if top[i] == OPEN and bottom[i] == ROUND_ROCK:
            return False

    return True


new_top, new_bottom = roll(platforms[0], platforms[1])

rolled_platforms = []

finished = [False] * len(platforms)
while not all(finished):
    rolled_platforms = []
    i = len(platforms) - 1
    bottom = platforms[i]
    while i > 0:
        new_top, new_bottom = roll(platforms[i-1], bottom)
        rolled_platforms.append(new_bottom)
        bottom = new_top
        i -= 1
    rolled_platforms.append(bottom)
    rolled_platforms.reverse()

    i = 1
    finished = [True]
    while i < len(platforms) - 2:
        finished.append(check_platforms(
            rolled_platforms[i], rolled_platforms[i + 1]))
        i += 1

    platforms = rolled_platforms


for platform in rolled_platforms:
    print(platform)

total = 0
i = len(rolled_platforms)
for j in range(len(rolled_platforms)):
    round_rocks = rolled_platforms[j].count(ROUND_ROCK)
    total += round_rocks * i
    i -= 1

print(total)
