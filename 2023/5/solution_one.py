import re
import sys

lines = []
with open("input.txt", "r") as file:
    lines = [line.rstrip() for line in file]

seeds = []

seeds_to_soil = []
soil_to_fert = []
fert_to_water = []
water_to_light = []
light_to_temp = []
temp_to_humidity = []
humidity_to_location = []


def parse_seeds(i):
    seeds_re = re.search(r"seeds: (.*)", lines[i])
    seeds = seeds_re.group(1).split(" ")
    return i + 1, seeds


def parse_ranges(i):
    ranges = []
    m = []
    j = i + 1
    while j < len(lines) and lines[j] != '':
        ranges.append(lines[j])
        j += 1

    print(ranges)
    for r in ranges:
        r = r.split(" ")
        dest, source, length = int(r[0]), int(r[1]), int(r[2])
        print(dest, source, length)
        source_start = source
        source_end = source + (length - 1)
        dest_start = dest
        dest_end = dest + (length - 1)
        m.append((source_start, source_end, dest_start, dest_end))

    return j, m


def find_dest(seed, map_range):
    for r in map_range:
        if seed >= r[0] and seed <= r[1]:
            return r[2] + (seed - r[0])
    return seed


i = 0
while i < len(lines):
    if lines[i] == '':
        i += 1
        continue

    if lines[i].startswith("seeds"):
        print('seeds')
        i, seeds = parse_seeds(i)
    elif lines[i].startswith("seed-to-soil"):
        print('seed-to-soil')
        i, seeds_to_soil = parse_ranges(i)
    elif lines[i].startswith("soil-to-fertilizer"):
        print('soil-to-fertilizer')
        i, soil_to_fert = parse_ranges(i)
    elif lines[i].startswith("fertilizer-to-water"):
        print('fertilizer-to-water')
        i, fert_to_water = parse_ranges(i)
    elif lines[i].startswith("water-to-light"):
        print('water-to-light')
        i, water_to_light = parse_ranges(i)
    elif lines[i].startswith("light-to-temp"):
        print('light-to-temp')
        i, light_to_temp = parse_ranges(i)
    elif lines[i].startswith("temperature-to-humidity"):
        print('temperature-to-humidity')
        i, temp_to_humidity = parse_ranges(i)
    elif lines[i].startswith("humidity-to-location"):
        print('humidity-to-location')
        i, humidity_to_location = parse_ranges(i)
    else:
        print('unknown')
        print(lines[i])


min_location = 9223372036854775807

for seed in seeds:
    soil, fert, water, light, temp, humidity, location = 0, 0, 0, 0, 0, 0, 0
    soil = find_dest(int(seed), seeds_to_soil)
    fert = find_dest(soil, soil_to_fert)
    water = find_dest(fert, fert_to_water)
    light = find_dest(water, water_to_light)
    temp = find_dest(light, light_to_temp)
    humidity = find_dest(temp, temp_to_humidity)
    location = find_dest(humidity, humidity_to_location)
    min_location = min(min_location, location)

print(min_location)
