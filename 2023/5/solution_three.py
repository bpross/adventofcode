import re

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


def find_source(dest, map_range):
    for r in map_range:
        if dest >= r[2] and dest <= r[3]:
            return r[0] + (dest - r[2])
    return dest


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

print(humidity_to_location)
min_location = 9223372036854775807
max_location = 0

seed_to_seed = {}
soil_to_seed = {}
fert_to_seed = {}
water_to_seed = {}
light_to_seed = {}
temp_to_seed = {}
humidity_to_seed = {}


searched = set()
ranges = 0
humidity_to_location = sorted(humidity_to_location, key=lambda x: x[3])
for r in humidity_to_location:
    ranges += 1
    loc_start, loc_end = r[2], r[3]
    print(ranges, loc_start, loc_end)
    while loc_start <= loc_end:
        if loc_start > min_location:
            break

        max_location = max(max_location, loc_start)
        if loc_start in searched:
            loc_start += 1
            continue

        print(loc_start, loc_end, min_location)
        searched.add(loc_start)
        humidity = find_source(loc_start, humidity_to_location)
        if humidity in humidity_to_seed:
            seed = humidity_to_seed[humidity]
            min_location = min(min_location, loc_start)
            print("found", seed, loc_start, min_location)
            break

        temp = find_source(humidity, temp_to_humidity)
        if temp in temp_to_seed:
            humidity_to_seed[humidity] = temp_to_seed[temp]
            seed = temp_to_seed[temp]
            min_location = min(min_location, loc_start)
            print("found", seed, loc_start, min_location)
            break

        light = find_source(temp, light_to_temp)
        if light in light_to_seed:
            humidity_to_seed[humidity] = light_to_seed[temp]
            temp_to_seed[temp] = light_to_seed[light]
            seed = light_to_seed[light]
            min_location = min(min_location, loc_start)
            print("found", seed, loc_start, min_location)
            break

        water = find_source(light, water_to_light)
        if water in water_to_seed:
            humidity_to_seed[humidity] = water_to_seed[temp]
            temp_to_seed[temp] = water_to_seed[water]
            light_to_seed[light] = water_to_seed[water]
            seed = water_to_seed[water]
            min_location = min(min_location, loc_start)
            print("found", seed, loc_start, min_location)
            break

        fert = find_source(water, fert_to_water)
        if fert in fert_to_seed:
            humidity_to_seed[humidity] = fert_to_seed[temp]
            temp_to_seed[temp] = fert_to_seed[water]
            light_to_seed[light] = fert_to_seed[water]
            water_to_seed[water] = fert_to_seed[fert]
            seed = fert_to_seed[fert]
            min_location = min(min_location, loc_start)
            print("found", seed, loc_start, min_location)
            break

        soil = find_source(fert, soil_to_fert)
        if soil in soil_to_seed:
            humidity_to_seed[humidity] = soil_to_seed[temp]
            temp_to_seed[temp] = soil_to_seed[water]
            light_to_seed[light] = soil_to_seed[water]
            water_to_seed[water] = soil_to_seed[fert]
            fert_to_seed[fert] = soil_to_seed[soil]
            seed = soil_to_seed[soil]
            min_location = min(min_location, loc_start)
            print("found", seed, loc_start, min_location)
            break

        seed = find_source(soil, seeds_to_soil)
        humidity_to_seed[humidity] = seed
        temp_to_seed[temp] = seed
        light_to_seed[light] = seed
        water_to_seed[water] = seed
        fert_to_seed[fert] = seed
        soil_to_seed[soil] = seed

        for seed_range in seeds:
            if seed >= seed_range[0] and seed <= seed_range[1]:
                min_location = min(min_location, loc_start)
                print("found", seed, loc_start, min_location)
                break

        loc_start += 1

loc_start, loc_end = 0, max_location
while loc_start <= loc_end:
    if loc_start in searched:
        loc_start += 1
        continue
    searched.add(loc_start)
    humidity = find_source(loc_start, humidity_to_location)
    temp = find_source(humidity, temp_to_humidity)
    light = find_source(temp, light_to_temp)
    water = find_source(light, water_to_light)
    fert = find_source(water, fert_to_water)
    soil = find_source(fert, soil_to_fert)
    seed = find_source(soil, seeds_to_soil)

    for seed_range in seeds:
        if seed >= seed_range[0] and seed <= seed_range[1]:
            min_location = min(min_location, loc_start)
            print("found", seed, loc_start, min_location)
            break

    loc_start += 1

print(min_location)
