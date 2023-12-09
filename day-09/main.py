import sys


def extrapolate(history, direction='forward'):
    data = [history]
    max = len(history)

    gen = 0

    for gen in range(0, max):
        tmp = []
        for i in range(len(data[gen])-1):
            tmp.append(data[gen][i+1] - data[gen][i])

        all_zero = [t == 0 for t in tmp]
        data.append(tmp)

        if all(all_zero):
            break

    if direction == 'forward':
        data[-1].append(0)

        for gen in range(len(data)-1, 0, -1):
            data[gen-1].append(data[gen-1][-1] + data[gen][-1])

        return data[0][-1]

    else:
        data[-1].insert(0, 0)

        for gen in range(len(data)-1, 0, -1):
            data[gen-1].insert(0, data[gen-1][0] - data[gen][0])

        return data[0][0]


file_name = sys.argv[1]
values = []

with open(file_name, 'r') as my_file:

    for line in my_file:
        history = line.strip().split(" ")
        history = [int(h) for h in history]
        values.append(history)

# Part 1
total = 0
for v in values:
    total += extrapolate(v, 'forward')

print("Part One")
print("Total: ", total)

# Part 2
total = 0
for v in values:
    total += extrapolate(v, 'backwards')

print("Part Two")
print("Total: ", total)
