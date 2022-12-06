def part1(datastream):
    s, position, length = set(), 0, 0
    for char in datastream:
        s.add(char)
        if not length < len(s):
            s = set()
        if len(s) == 4:
            return position
        length = len(s)
        position += 1

input = open("input.txt", "r")
datastream = input.readline()

print("Solution to part one =>", part1(datastream))
