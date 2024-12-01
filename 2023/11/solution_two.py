total_expansion = 999999

universe = []
with open('input.txt', 'r') as f:
    for line in f:
        universe.append(list(line.strip()))

rows_to_expand = []
cols_to_expand = []
for i in range(len(universe)):
    if '#' not in universe[i]:
        rows_to_expand.append(i)

for i in range(len(universe[0])):
    if '#' not in [row[i] for row in universe]:
        cols_to_expand.append(i)

# print(rows_to_expand)
# print(cols_to_expand)
# rows_to_expand.reverse()
# cols_to_expand.reverse()


# for i in cols_to_expand:
#     for row in universe:
#         for _ in range(total_expansion):
#             row.insert(i, ".")

# row_expansion = ["."] * len(universe[0])
# for i in rows_to_expand:
#     for _ in range(total_expansion):
#         universe.insert(i + 1, row_expansion.copy())


galaxies = []
for i in range(len(universe)):
    for j in range(len(universe[i])):
        if universe[i][j] == '#':
            r, c = i, j
            for row in rows_to_expand:
                if i > row:
                    r += total_expansion
            for col in cols_to_expand:
                if j > col:
                    c += total_expansion
            galaxies.append((r, c))

total_distances = 0
found = 0
for i in range(len(galaxies)):
    for j in range(i + 1, len(galaxies)):
        found += 1
        total_distances += abs(galaxies[j][0] - galaxies[i][0]) + \
            abs(galaxies[j][1] - galaxies[i][1])

print(found)
print(total_distances)
