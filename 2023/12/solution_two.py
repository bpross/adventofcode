lines = []
with open('sample.txt') as f:
    for line in f:
        lines.append(line.strip().split(" "))


def validate(springs, groups):
    stack = []
    itr = 0
    while itr < len(springs):
        group = ""
        while itr < len(springs) and (springs[itr] == "#" or springs[itr] == "?"):
            group += "#"
            itr += 1
        if group != "":
            stack.append(group)

        while itr < len(springs) and springs[itr] == ".":
            itr += 1
    if len(stack) != len(groups):
        return False

    while len(groups) > 0 and len(stack) > 0:
        group = groups.pop()
        spring_group = stack.pop()
        if "?" in spring_group:
            return False
        if int(group) != len(spring_group):
            return False

    return True


def generate(springs, itr, groups):
    if itr == len(springs):
        g = groups.copy()
        return 1 if validate(springs, g) else 0

    s_copy = springs.copy()
    period_copy = s_copy.copy()
    pound_copy = s_copy.copy()
    total = 0
    if s_copy[itr] == "." or s_copy[itr] == "#":
        total += generate(s_copy, itr + 1, groups)
    else:
        period_copy[itr] = "."
        pound_copy[itr] = "#"
        total += generate(period_copy, itr + 1, groups)
        total += generate(pound_copy, itr + 1, groups)

    return total


total = 0
checked = 0
for line in lines:
    org_groups = line[1].split(",")
    springs = line[0]
    org_springs = [*springs]
    groups, springs = [], []
    for i in range(5):
        groups.extend(org_groups)
        springs.extend(org_springs)
        springs.append("?")

    total += generate(springs, 0, groups)
    checked += 1
    print(checked)

print(total)
