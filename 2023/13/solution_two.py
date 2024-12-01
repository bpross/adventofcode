from copy import deepcopy

patterns = []

with open('sample.txt') as f:
    pattern = []
    for line in f:
        if line == "\n":
            patterns.append(pattern)
            pattern = []
            continue

        pattern.append([c for c in line.strip()])

    # get the last one
    patterns.append(pattern)


def transpose(pattern):
    return list(map(list, zip(*pattern)))


def find_reflection_point_horizontal(pattern):
    start, end = 0, 1
    found = False

    while end < len(pattern):
        if pattern[start] == pattern[end]:
            # need to expand out, its a valid pattern if either check_start or check_end
            # over/underflows
            check_start, check_end = start - 1, end + 1
            while check_start >= 0 and check_end < len(pattern):
                if pattern[check_start] == pattern[check_end]:
                    check_start -= 1
                    check_end += 1
                else:
                    break

            found = True if check_start < 0 or check_end >= len(
                pattern) else False
            if found:
                break

        start += 1
        end += 1

    return found, start


def find_original_reflection_point(pattern):
    horizontal, horizontal_start = find_reflection_point_horizontal(
        pattern)
    if horizontal:
        return True, horizontal_start + 1, False, - 1
    vertcial, vertical_start = find_reflection_point_horizontal(
        transpose(pattern))
    if vertcial:
        return False, -1, True, vertical_start + 1


total = 0
p = 0

for pattern in patterns:
    org_horizontal, org_horizontal_start, org_vertical, org_vertical_start = find_original_reflection_point(
        pattern)

    smudge = False
    for i in range(len(pattern)):
        for j in range(len(pattern[i])):
            p_copy = deepcopy(pattern)
            if p_copy[i][j] == ".":
                p_copy[i][j] = "#"
            else:
                p_copy[i][j] = "."

            for row in p_copy:
                print(''.join(row))
            print("\n")

            horizontal, horizontal_start = find_reflection_point_horizontal(
                p_copy)
            if horizontal:
                if org_horizontal and horizontal_start + 1 == org_horizontal_start:
                    continue
                else:
                    total += (horizontal_start + 1) * 100
                    p += 1
                    smudge = True
                    break
            vertcial, vertical_start = find_reflection_point_horizontal(
                transpose(p_copy))
            if vertcial:
                if org_vertical and vertical_start + 1 == org_vertical_start:
                    continue
                else:
                    total += vertical_start + 1
                    smudge = True
                    break

        if smudge:
            break

    if not smudge:
        print("smudge not found", p)

    p += 1

print("total", total)
