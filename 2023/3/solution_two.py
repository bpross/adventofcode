f = open('input.txt', 'r')

engine = []

for line in f:
    number = ''
    engine_line = []
    for c in line:
        if c == '\n':
            continue
        engine_line.append(c)

    engine.append(engine_line)

print(engine)

total = 0
visited = set()

def get_cords(x, y):
    cords = []
    if x - 1 >= 0 and y - 1 >= 0:
        cords.append((x - 1, y - 1))
    if x - 1 >= 0 and y >= 0:
        cords.append((x - 1, y))
    if x - 1 >= 0 and y + 1 < len(engine):
        cords.append((x - 1, y + 1))
    if x >= 0 and y + 1 < len(engine):
        cords.append((x, y + 1))
    if x + 1 < len(engine[y]) and y + 1 < len(engine):
        cords.append((x + 1, y + 1))
    if x + 1 < len(engine[y]) and y >= 0:
        cords.append((x + 1, y))
    if x + 1 < len(engine[y]) and y - 1 >= 0:
        cords.append((x + 1, y - 1))
    if x >= 0 and y - 1 >= 0:
        cords.append((x, y - 1))

    return cords

for y in range(0, len(engine)):
    for x in range(0, len(engine[y])):
        if engine[y][x] != '*' or engine[y][x].isnumeric():
            continue
        cords = get_cords(x,y)
        ratios = []
        for cord in cords:
            if cord in visited:
                continue
            if engine[cord[1]][cord[0]].isnumeric():
                print(f'found: {engine[cord[1]][cord[0]]}')
                # consume left and right to construct the number
                number = engine[cord[1]][cord[0]]
                # go right first
                for i in range(cord[0] + 1, len(engine[cord[1]])):
                    if not engine[cord[1]][i].isdigit():
                        break
                    number += engine[cord[1]][i]
                    visited.add((i, cord[1]))
                print(f'right: {number}')
                # go left
                left = ''
                for i in range(cord[0] - 1, -1, -1):
                    if not engine[cord[1]][i].isdigit():
                        break
                    left = engine[cord[1]][i] + left 
                    visited.add((i, cord[1]))
                number = left + number
                print(f'final: {number}')
                ratios.append(number)
            visited.add(cord)

        if len(ratios) == 2:
            total += int(ratios[0]) * int(ratios[1])

print(total)
