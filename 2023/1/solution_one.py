f = open('input', 'r')
total = 0
for line in f:
    first_found = False
    first = 0
    last = 0
    for c in line:
        if c.isdigit():
            if not first_found:
                first = int(c)
                first_found = True
            last = int(c)
    first_found = False
    total += first * 10 + last

print(total)
