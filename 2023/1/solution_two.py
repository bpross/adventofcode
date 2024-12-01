f = open('input_two.txt', 'r')
total = 0

numbers = {
    'one': '1',
    'two': '2',
    'three': '3',
    'four': '4',
    'five': '5',
    'six': '6',
    'seven': '7',
    'eight': '8',
    'nine': '9'
}

modified_lines = []

for line in f:
    i = 0
    new_line = ''
    while i < len(line):
        found = False
        for key in numbers:
            print(line[i:])
            if line[i:].startswith(key):
                print('found')
                new_line += numbers[key]
                i += len(key) - 1
                found = True
                break
            else:
                new_line += line[i]
        if not found:
            i += 1
            found = False
    modified_lines.append(new_line)

print(modified_lines)

for line in modified_lines:
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
