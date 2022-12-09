package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SIZE = 1000

type Grid [SIZE][SIZE]string

func (grid Grid) printGrid() {
	for i := range grid {
		fmt.Println(grid[i])
	}
}

type Coordinates struct {
	x, y int
}

func checkVisited(square *string, res *int) {
	if *square != "#" {
		*square = "#"
		*res++
	}
}

func compare(coordH, coordT *Coordinates) {
	if coordT.x != coordH.x {
		coordT.x = coordH.x
	}
}

func part1() int {
	var (
		grid     Grid
		input, _ = os.Open("input.txt")
		reader   = bufio.NewReader(input)
		startX   = SIZE / 2
		startY   = SIZE / 2
		coordH   = Coordinates{startX, startY}
		coordT   = coordH
		result   = 1
	)
	grid[coordT.x][coordT.y] = "#"
	grid[coordH.x][coordH.y] = "#"
	for {
		line, _ := reader.ReadString('\n')
		if line == "" {
			break
		}
		line = strings.Trim(line, "\n")
		var (
			args    = strings.Split(line, " ")
			iter, _ = strconv.Atoi(args[1])
		)
		switch args[0] {
		case "R":
			for ; iter > 0; iter-- {
				coordH.y++
				if coordH.y-coordT.y > 1 {
					coordT.y++
					compare(&coordH, &coordT)
				}
				checkVisited(&grid[coordT.x][coordT.y], &result)
			}
		case "L":
			for ; iter > 0; iter-- {
				coordH.y--
				if coordT.y-coordH.y > 1 {
					coordT.y--
					compare(&coordH, &coordT)
				}
				checkVisited(&grid[coordT.x][coordT.y], &result)
			}
		case "U":
			for ; iter > 0; iter-- {
				coordH.x--
				if coordT.x-coordH.x > 1 {
					coordT.x--
					compare(&coordH, &coordT)
				}
				checkVisited(&grid[coordT.x][coordT.y], &result)
			}
		case "D":
			for ; iter > 0; iter-- {
				coordH.x++
				if coordH.x-coordT.x > 1 {
					coordT.x++
					compare(&coordH, &coordT)
				}
				checkVisited(&grid[coordT.x][coordT.y], &result)
			}
		}
	}
	return result
}

func part2() int {
	return 0
}

func main() {
	fmt.Println("Solution to part one =>", part1())
	fmt.Println("Solution to part two =>", part2())
}
