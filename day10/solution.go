package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkCycle(cycle, X int) int {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		return cycle * X
	}
	return 0
}

func part1() int {
	var (
		X, cycle, sum = 1, 0, 0
		input, _      = os.Open("input.txt")
		reader        = bufio.NewReader(input)
	)
	for {
		line, _ := reader.ReadString('\n')
		if line == "" {
			break
		}
		line = strings.Trim(line, "\n")
		if line == "noop" {
			cycle++
			sum += checkCycle(cycle, X)
		} else {
			V, _ := strconv.Atoi(strings.Split(line, " ")[1])
			cycle++
			sum += checkCycle(cycle, X)
			cycle++
			sum += checkCycle(cycle, X)
			X += V
		}
	}
	return sum
}

func main() {
	fmt.Println("Solution to part one =>", part1())
}
