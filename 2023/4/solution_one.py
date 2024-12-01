import re

f = open("input.txt", "r")

total = 0

for card in f:
    card_re = re.search(r"Card(\s+)(\d+): (.+)", card)
    numbers = card_re.group(3).split("|")
    winning_numbers = numbers[0].strip().split(" ")
    our_numbers = numbers[1].strip().split(" ")

    winners = {}
    for number in winning_numbers:
        if number == '':
            continue
        winners[number] = True

    total_winners = 0
    for number in our_numbers:
        if number == '':
            continue
        if number in winners:
            total_winners += 1

    if total_winners > 0:
        total += 2 ** (total_winners - 1)

print(total)
