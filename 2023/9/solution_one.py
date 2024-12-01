lines = []

with open("input.txt", "r") as f:
    for line in f:
        lines.append(line.strip().split(" "))

differences = []

for line in lines:
    print(line)
    all_zeros = False 
    check_line = line
    diffs = [line]
    while not all_zeros:
        diff = []
        for i in range(len(check_line)):
            if i + 1 == len(check_line):
                break
            val = int(check_line[i + 1]) - int(check_line[i])
            diff.append(val)

        for val in diff:
            if val != 0:
                all_zeros = False
                break
            all_zeros = True
        if not diff:
            diff = [0]
            all_zeros = True

        check_line = diff
        diffs.append(diff)

    differences.append(diffs)

totals = []
for row_diffs in differences:
    total = 0
    for i, diff in reversed(list(enumerate(row_diffs))):
        total = total + int(diff[-1])

    totals.append(total)

print(sum(totals))
