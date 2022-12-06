def part1(datastream):
    marker, position, length = set(), 0, 0
    for char in datastream:
        marker.add(char)
        if not length < len(s):
            marker.clear() 
        elif len(marker) == 4:
            return position
        length = len(marker)
        position += 1

def part2(datastream):
    marker, position, = [], 0
    for char in datastream:
        if char in marker:
            marker = marker[marker.index(char)+1:]
        marker.append(char)
        position += 1
        if len(marker) == 14:
            return position

input = open("input.txt", "r")
datastream = input.readline()
print(f"Solution to part one => {part1(datastream)}")
print(f"Solution to part two => {part2(datastream)}")
