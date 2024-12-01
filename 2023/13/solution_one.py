patterns = []

with open('input.txt') as f:
    pattern = []
    for line in f:
        if line == "\n":
            patterns.append(pattern)
            pattern = []
            continue

        pattern.append(line.strip())

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


total = 0
p = 0

for pattern in patterns:
    horizontal, horizontal_start = find_reflection_point_horizontal(pattern)
    if horizontal:
        total += (horizontal_start + 1) * 100
        p += 1
        continue
    vertcial, vertical_start = find_reflection_point_horizontal(
        transpose(pattern))
    if vertcial:
        total += vertical_start + 1

    p += 1

print("total", total)
