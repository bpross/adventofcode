import re
from collections import OrderedDict, defaultdict

REMOVE = "-"
ADD = "="


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

boxes = defaultdict(OrderedDict)

for step in steps:
    label_search = re.search(r"(\w{1,6})(=(\d{1,3})|-)", step)
    label = label_search.group(1)
    box_number = hash(label)
    operation = label_search.group(2)
    if operation == REMOVE:
        if label in boxes[box_number]:
            del boxes[box_number][label]

    else:
        # Add
        boxes[box_number][label] = label_search.group(3)


print(boxes)
total = 0
for box in boxes:
    slot = 1
    for label in boxes[box]:
        box_slot = (box + 1) * slot
        focal_slot = int(boxes[box][label])
        focus_power = box_slot * focal_slot
        total += focus_power
        slot += 1

print(total)
