import re

lines = []
with open("sample.txt", "r") as file:
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
    seed_ranges = seeds_re.group(1).split(" ")
    k = 0
    while k < len(seed_ranges):
        start = int(seed_ranges[k])
        num_seeds = int(seed_ranges[k + 1])
        seeds.append((start, start + num_seeds))
        k += 2

    return i + 1, seeds


def parse_ranges(i):
    ranges = []
    m = []
    j = i + 1
    while j < len(lines) and lines[j] != '':
        ranges.append(lines[j])
        j += 1

    for r in ranges:
        r = r.split(" ")
        dest, source, length = int(r[0]), int(r[1]), int(r[2])
        source_start = source
        source_end = source + (length - 1)
        dest_start = dest
        dest_end = dest + (length - 1)
        m.append((source_start, source_end, dest_start, dest_end))

    return j, m


def find_dest(seed_range, map_range):
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
        i, seeds = parse_seeds(i)
    elif lines[i].startswith("seed-to-soil"):
        i, seeds_to_soil = parse_ranges(i)
    elif lines[i].startswith("soil-to-fertilizer"):
        i, soil_to_fert = parse_ranges(i)
    elif lines[i].startswith("fertilizer-to-water"):
        i, fert_to_water = parse_ranges(i)
    elif lines[i].startswith("water-to-light"):
        i, water_to_light = parse_ranges(i)
    elif lines[i].startswith("light-to-temp"):
        i, light_to_temp = parse_ranges(i)
    elif lines[i].startswith("temperature-to-humidity"):
        i, temp_to_humidity = parse_ranges(i)
    elif lines[i].startswith("humidity-to-location"):
        i, humidity_to_location = parse_ranges(i)
    else:
        print('unknown')
        print(lines[i])

print(seeds)
print(seeds_to_soil)

min_location = 9223372036854775807

seed_to_dest = {}
soil_to_dest = {}
fert_to_dest = {}
water_to_dest = {}
light_to_dest = {}
temp_to_dest = {}
humidity_to_dest = {}


range_num = 0
for seed_range in seeds:
    range_num += 1
    start, num_seeds = seed_range
    for seed in range(start, start + num_seeds):
        soil, fert, water, light, temp, humidity, location = 0, 0, 0, 0, 0, 0, 0
        if seed in seed_to_dest:
            min_location = min(min_location, seed_to_dest[seed])
            continue
        else:
            soil = find_dest(seed_range, seeds_to_soil)

        if soil in soil_to_dest:
            seed_to_dest[seed] = soil_to_dest[soil]
            min_location = min(min_location, soil_to_dest[soil])
            continue
        else:
            fert = find_dest(soil, soil_to_fert)

        if fert in fert_to_dest:
            seed_to_dest[seed] = soil_to_dest[soil]
            soil_to_dest[soil] = fert_to_dest[fert]
            min_location = min(min_location, fert_to_dest[fert])
            continue
        else:
            water = find_dest(fert, fert_to_water)

        if water in water_to_dest:
            seed_to_dest[seed] = soil_to_dest[soil]
            soil_to_dest[soil] = fert_to_dest[fert]
            fert_to_dest[fert] = water_to_dest[water]
            min_location = min(min_location, water_to_dest[water])
            continue
        else:
            light = find_dest(water, water_to_light)

        if light in light_to_dest:
            seed_to_dest[seed] = soil_to_dest[soil]
            soil_to_dest[soil] = fert_to_dest[fert]
            fert_to_dest[fert] = water_to_dest[water]
            water_to_dest[water] = light_to_dest[light]
            min_location = min(min_location, light_to_dest[light])
            continue
        else:
            temp = find_dest(light, light_to_temp)

        if temp in temp_to_dest:
            seed_to_dest[seed] = soil_to_dest[soil]
            soil_to_dest[soil] = fert_to_dest[fert]
            fert_to_dest[fert] = water_to_dest[water]
            water_to_dest[water] = light_to_dest[light]
            light_to_dest[light] = temp_to_dest[temp]
            min_location = min(min_location, temp_to_dest[temp])
            continue
        else:
            humidity = find_dest(temp, temp_to_humidity)

        if humidity in humidity_to_dest:
            seed_to_dest[seed] = soil_to_dest[soil]
            soil_to_dest[soil] = fert_to_dest[fert]
            fert_to_dest[fert] = water_to_dest[water]
            water_to_dest[water] = light_to_dest[light]
            light_to_dest[light] = temp_to_dest[temp]
            temp_to_dest[temp] = humidity_to_dest[humidity]
            min_location = min(min_location, humidity_to_dest[humidity])
            continue
        else:
            location = find_dest(humidity, humidity_to_location)

        humidity_to_dest[humidity] = location
        temp_to_dest[temp] = humidity_to_dest[humidity]
        light_to_dest[light] = temp_to_dest[temp]
        water_to_dest[water] = light_to_dest[light]
        fert_to_dest[fert] = water_to_dest[water]
        soil_to_dest[soil] = fert_to_dest[fert]
        seed_to_dest[seed] = soil_to_dest[soil]
        min_location = min(min_location, location)
        print(range_num, seed, soil, fert, water,
              light, temp, humidity, location)

print(min_location)
