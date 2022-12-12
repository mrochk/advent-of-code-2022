package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const size = 1000

type Grid [size][size]string

type Coordinates struct {
	x, y int
}

func checkVisited(square *string, res *int) {
	if *square != "#" {
		*square = "#"
		*res++
	}
}

func cmp(x bool, coordH, coordT *Coordinates) {
	if x {
		if coordT.x != coordH.x {
			coordT.x = coordH.x
		}
	} else {

		if coordT.y != coordH.y {
			coordT.y = coordH.y
		}
	}
}

func part1() int {
	var (
		grid     Grid
		input, _ = os.Open("input.txt")
		reader   = bufio.NewReader(input)
		startX   = size / 2
		startY   = startX
		head     = Coordinates{startX, startY}
		tail     = head
		result   = 1
	)
	grid[head.x][head.y] = "#"
	for {
		line, _ := reader.ReadString('\n')
		if line == "" {
			break
		}
		line = strings.Trim(line, "\n")
		args := strings.Split(line, " ")
		iter, _ := strconv.Atoi(args[1])
		switch args[0] {
		case "R":
			for ; iter > 0; iter-- {
				head.y++
				if head.y-tail.y > 1 {
					tail.y++
					cmp(true, &head, &tail)
				}
				checkVisited(&grid[tail.x][tail.y], &result)
			}
		case "L":
			for ; iter > 0; iter-- {
				head.y--
				if tail.y-head.y > 1 {
					tail.y--
					cmp(true, &head, &tail)
				}
				checkVisited(&grid[tail.x][tail.y], &result)
			}
		case "U":
			for ; iter > 0; iter-- {
				head.x--
				if tail.x-head.x > 1 {
					tail.x--
					cmp(false, &head, &tail)
				}
				checkVisited(&grid[tail.x][tail.y], &result)
			}
		case "D":
			for ; iter > 0; iter-- {
				head.x++
				if head.x-tail.x > 1 {
					tail.x++
					cmp(false, &head, &tail)
				}
				checkVisited(&grid[tail.x][tail.y], &result)
			}
		}
	}
	return result
}

func main() {
	fmt.Println("Solution to part one =>", part1())
}
