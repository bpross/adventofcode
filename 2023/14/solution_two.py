ROUND_ROCK = "O"
SQUARE_ROCK = "#"
OPEN = "."

CYCLES = 1000

platforms = []

with open('input.txt') as f:
    for line in f:
        platforms.append([c for c in line.strip()])


def roll_north(top, bottom):
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


def roll_south(top, bottom):
    new_top, new_bottom = [], []
    for i in range(len(top)):
        if bottom[i] == ROUND_ROCK or bottom[i] == SQUARE_ROCK or (bottom[i] == OPEN and top[i] == OPEN) or (bottom[i] == OPEN and top[i] == SQUARE_ROCK):
            new_top.append(top[i])
            new_bottom.append(bottom[i])
        elif bottom[i] == OPEN and top[i] == ROUND_ROCK:
            new_top.append(OPEN)
            new_bottom.append(top[i])
        else:
            print(top[i], bottom[i])
    return new_top, new_bottom


def roll_west(row):
    i = 0
    while i < len(row):
        if row[i] == OPEN and i + 1 < len(row) and row[i + 1] == ROUND_ROCK:
            row[i], row[i + 1] = row[i + 1], row[i]

        i += 1

    return row


def roll_east(row):
    i = len(row) - 1
    while i >= 0:
        if row[i] == OPEN and i - 1 >= 0 and row[i - 1] == ROUND_ROCK:
            row[i], row[i - 1] = row[i - 1], row[i]
        i -= 1
    return row


def check_north(top, bottom):
    for i in range(len(top)):
        if top[i] == OPEN and bottom[i] == ROUND_ROCK:
            return False

    return True


def check_south(top, bottom):
    for i in range(len(top)):
        if bottom[i] == OPEN and top[i] == ROUND_ROCK:
            return False
    return True


def check_west(row):
    i = 0
    while i < len(row):
        if row[i] == OPEN and i + 1 < len(row) and row[i + 1] == ROUND_ROCK:
            return False
        i += 1
    return True


def check_east(row):
    i = len(row) - 1
    while i >= 0:
        if row[i] == OPEN and i - 1 >= 0 and row[i - 1] == ROUND_ROCK:
            return False
        i -= 1
    return True


def print_platforms(platforms):
    for platform in platforms:
        print(platform)

    print("\n")


seen = {}

for k in range(CYCLES):
    print(k)
    # NORTH
    finished = [False] * len(platforms)
    while not all(finished):
        rolled_platforms = []
        i = len(platforms) - 1
        bottom = platforms[i]
        while i > 0:
            new_top, new_bottom = roll_north(platforms[i-1], bottom)
            rolled_platforms.append(new_bottom)
            bottom = new_top
            i -= 1
        rolled_platforms.append(bottom)
        rolled_platforms.reverse()

        i = 1
        finished = [True]
        while i < len(platforms) - 2:
            finished.append(check_north(
                rolled_platforms[i], rolled_platforms[i + 1]))
            i += 1

        platforms = rolled_platforms

    # WEST
    for platform in platforms:
        while not check_west(platform):
            platform = roll_west(platform)

    # SOUTH
    finished = [False] * len(platforms)
    while not all(finished):
        rolled_platforms = []
        i = len(platforms) - 1
        bottom = platforms[i]
        while i > 0:
            new_top, new_bottom = roll_south(platforms[i - 1], bottom)
            rolled_platforms.insert(0, new_bottom)
            bottom = new_top
            i -= 1
        rolled_platforms.insert(0, bottom)

        i = 1
        finished = [True]
        while i < len(platforms) - 1:
            finished.append(check_south(
                rolled_platforms[i], rolled_platforms[i + 1]))
            i += 1

        platforms = rolled_platforms

    # EAST
    for platform in platforms:
        while not check_east(platform):
            platform = roll_east(platform)

    s = ""
    for platform in platforms:
        s += "".join(platform)
    h = hash(s)
    if h in seen:
        print("Found cycle at", k)
        print("Cycle length", k - seen[h])
        break
    else:
        seen[h] = k

for platform in platforms:
    print(platform)

total = 0
i = len(platforms)
for j in range(len(platforms)):
    round_rocks = platforms[j].count(ROUND_ROCK)
    total += round_rocks * i
    i -= 1

print(total)
