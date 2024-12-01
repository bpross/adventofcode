import re
from collections import defaultdict

f = open("input.txt", "r")

cards = defaultdict(int)
total_wins = defaultdict(int)

for card in f:
    card_re = re.search(r"Card(\s+)(\d+): (.+)", card)
    id = int(card_re.group(2))
    cards[id] += 1

    numbers = card_re.group(3).split("|")
    winning_numbers = numbers[0].strip().split(" ")
    our_numbers = numbers[1].strip().split(" ")

    winners = {}
    for number in winning_numbers:
        if number == '':
            continue
        winners[number] = True

    for number in our_numbers:
        if number == '':
            continue
        if number in winners:
            total_wins[id] += 1


# for key in total_wins.keys():
#     value = total_wins[key]
#     print(key, value)
#     print(f'kv: {key} {value}')
#     if value > 0:
#         for i in range(1, value + 1):
#             print(key + i, cards[key+i])
#             print(f'incrementing: {key+i}, has: {cards[key+i]})
#             cards[key + i] += 1

for id in total_wins.keys():
    wins = total_wins[id]
    if wins > 0:
        for c in range(cards[id]):
            for i in range(1, wins + 1):
                cards[id + i] += 1

print(total_wins)
print(cards)
print(sum(cards.values()))
