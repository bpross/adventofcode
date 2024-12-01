from collections import defaultdict

card_rank = 'J23456789TQKA'

hands = []
five_of_a_kind = []
four_of_a_kind = []
full_house = []
three_of_a_kind = []
two_pairs = []
one_pair = []
high_card = []

with open('input.txt') as f:
    for line in f:
        hands.append(line.strip().split(' '))


for hand in hands:
    counts = defaultdict(int)
    jokers = 0
    for card in hand[0]:
        if card == 'J':
            jokers += 1
        else:
            counts[card] += 1

    pairs, threes, fours, fives = 0, 0, 0, 0
    for c in card_rank:
        if counts[c] == 2:
            pairs += 1
        elif counts[c] == 3:
            threes += 1
        elif counts[c] == 4:
            fours += 1
        elif counts[c] == 5:
            fives += 1

    # print(hand, pairs, threes, fours, fives, jokers)
    # Five jokers, five of a kind, four of a kind +1 joker, three of a kind + 2 jokers, 1 pair + 3 jokers, four jokers
    if jokers == 5 or fives == 1 or (fours == 1 and jokers == 1) or (threes == 1 and jokers == 2) or (pairs == 1 and jokers == 3) or jokers == 4:
        five_of_a_kind.append(hand)
    # Four of a kind, three of a kind and 1 joker, one pair and two jokers
    elif fours == 1 or (threes == 1 and jokers == 1) or (pairs == 1 and jokers == 2) or jokers == 3:
        four_of_a_kind.append(hand)
    # 1 threes and 1 pair, 2 pairs and 1 joker
    elif threes == 1 and pairs == 1 or (pairs == 2 and jokers == 1):
        full_house.append(hand)
    # threes == 1, pairs == 1 and jokers = 1
    elif threes == 1 or (pairs == 1 and jokers == 1) or jokers == 2:
        three_of_a_kind.append(hand)
    # 2 pairs, 1 pair and 1 joker
    elif pairs == 2 or (pairs == 1 and jokers == 1):
        two_pairs.append(hand)
    # 1 pair, 1 joker
    elif pairs == 1 or jokers == 1:
        one_pair.append(hand)
    else:
        high_card.append(hand)


five_of_a_kind.sort(
    key=lambda hand: [card_rank.index(c) for c in hand[0]])
four_of_a_kind.sort(
    key=lambda hand: [card_rank.index(c) for c in hand[0]])
full_house.sort(key=lambda hand: [card_rank.index(c) for c in hand[0]])
three_of_a_kind.sort(key=lambda hand: [card_rank.index(c) for c in hand[0]])
two_pairs.sort(key=lambda hand: [card_rank.index(c) for c in hand[0]])
one_pair.sort(key=lambda hand: [card_rank.index(c) for c in hand[0]])
high_card.sort(key=lambda hand: [card_rank.index(c) for c in hand[0]])

ranked_hands = high_card + one_pair + two_pairs + \
    three_of_a_kind + full_house + four_of_a_kind + five_of_a_kind

total = 0
hand = 1
for h in ranked_hands:
    total += int(h[1]) * hand
    hand += 1

print(total)
