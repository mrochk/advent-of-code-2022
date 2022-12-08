input = open("input.txt", "r")
grid = []
while True:
    line = input.readline()
    if line == "": break
    grid.append(line.strip())
visible_trees, max_scenic_score, line_index = 0, 0, -1
for line in grid:
    line_index += 1
    column_index = -1
    for number in line:
        column_index += 1
        if line_index in [0, len(grid)-1] or column_index in [0, len(line)-1]:
            visible_trees += 1
        else:
            is_tree_visible = False
            visible, vdist_left = True, 0
            for i in range(column_index-1, -1, -1):
                vdist_left += 1
                if int(line[i]) >= int(number):
                    visible = False
                    break
            if visible: is_tree_visible = True
            visible, vdist_right = True, 0
            for i in range(column_index+1, len(line)):
                vdist_right += 1
                if int(line[i]) >= int(number):
                    visible = False
                    break
            if visible: is_tree_visible = True
            visible, vdist_top = True, 0
            for i in range(line_index-1, -1, -1):
                vdist_top += 1
                if int(grid[i][column_index]) >= int(number):
                    visible = False
                    break
            if visible: is_tree_visible = True
            visible, vdist_bottom = True, 0
            for i in range(line_index+1, len(grid)):
                vdist_bottom += 1
                if int(grid[i][column_index]) >= int(number):
                    visible = False
                    break
            if visible: is_tree_visible = True
            if is_tree_visible: visible_trees += 1
            scenic_score = vdist_left * vdist_right * vdist_bottom * vdist_top
            if scenic_score > max_scenic_score: max_scenic_score = scenic_score 

print(f"Solution to part one => {visible_trees}")
print(f"Solution to part two => {max_scenic_score}")