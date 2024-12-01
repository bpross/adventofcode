import re

f = open('input_one.txt', 'r')
red = 12
green = 13
blue = 14

total = 0

for game in f:
    max_red, max_green, max_blue = 0,0,0
    group_re = re.search(r"Game (\d+): (.+)", game)
    game_id = group_re.group(1)
    grabs = group_re.group(2).split(";")
    for grab in grabs:
        red_re = re.search(r"(\d+) red", grab)
        green_re = re.search(r"(\d+) green", grab)
        blue_re = re.search(r"(\d+) blue", grab)
        if red_re and red_re.lastindex:
            found_red = int(red_re.group(1))
            max_red = max(max_red, found_red)
        if green_re and green_re.lastindex:
            found_green = int(green_re.group(1))
            max_green = max(max_green, found_green)
        if blue_re and blue_re.lastindex:
            found_blue= int(blue_re.group(1))
            max_blue = max(max_blue, found_blue)

    if max_red <= red and max_green <= green and max_blue <= blue:
        total += int(game_id)
    
print(total)
