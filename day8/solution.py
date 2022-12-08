input = open("input.txt", "r")

# ligne, colonne
grid = []
while True:
    line = input.readline().strip()
    if line == "":
        break
    grid.append(line)

visible_trees = 0
visbl = []
line_index = -1
for line in grid:
    line_index += 1
    column_index = -1
    for number in line:
        column_index += 1
        if line_index in [0, len(grid)-1] or column_index in [0, len(line)-1]:
            visible_trees += 1
            visbl.append(number)
        else:
            X = False
            
            #check left
            visible = True
            for i in range(column_index-1, -1, -1):
                if int(line[i]) >= int(number):
                    visible = False
                    break

            if visible:
                X = True

            #check right
            visible = True
            for i in range(column_index+1, len(line)):
                if int(line[i]) >= int(number):
                    visible = False
                    break

            if visible:
                X = True

            #check top
            visible = True
            for i in range(line_index-1, -1, -1):
                if number == "1":
                    print(f"checking {grid[i][column_index]}")
                if int(grid[i][column_index]) >= int(number):
                    visible = False
                    
            if visible:
                X = True

            #check bottom
            visible = True
            for i in range(line_index+1, len(grid)):
                if int(grid[i][column_index]) >= int(number):
                    visible = False

            if visible:
                X = True
            
            if X:
                visible_trees += 1
                visbl.append(number)

print(visbl)
print(f"Solution to part one => {visible_trees}")