import sys
from math import lcm


def finished1(current):
    if current == 'ZZZ':
        return True


def finished2(current):
    if current[2] == 'Z':
        return True


def compute(current, finished):
    arrive = False
    count = 0
    i = 0

    while arrive is False:
        dir = instructions[i]
        direction = int(dir)
        current = map[current][direction]
        count += 1
        i += 1
        if finished(current):
            break
        if i == len(instructions):
            i = 0

    return count


file_name = sys.argv[1]
map = {}

with open(file_name, 'r') as my_file:

    instructions = my_file.readline().strip().replace('L', '0').replace('R', '1')

    for line in my_file:
        my_line = line.strip().replace(" ", "")
        if len(my_line) == 0:
            continue

        key, values_string = my_line.split("=")
        values = values_string.replace('(', '').replace(')', '').split(',')

        map[key] = values

print(instructions)

# Part 1
start = 'AAA'
count = compute(start, finished1)

print("Part One")
print("Count: ", count)

# Part 2
start_nodes = [k for k in map.keys() if k[2] == 'A']
steps = [compute(start, finished2) for start in start_nodes]

print("Part Two")
print("Count: ", lcm(*steps))
