def hash(s):
    current_value = 0
    for char in s:
        ascii_code = ord(char)
        current_value += ascii_code
        current_value *= 17
        current_value = current_value % 256

    return current_value


steps = []
with open('input.txt', 'r') as f:
    for line in f:
        for step in line.strip().split(','):
            steps.append(step)

total = 0
for step in steps:
    total += hash(step)

print(total)
