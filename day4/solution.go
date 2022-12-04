package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() int {
	var (
		input, _ = os.Open("input.txt")
		reader   = bufio.NewReader(input)
		total    = 0
		line     string
		args     []string
		intArgs  [4]int
	)
	for {
		line, _ = reader.ReadString('\n')
		if line == "" {
			break
		}
		line = strings.ReplaceAll(line, "-", ",")
		line = strings.ReplaceAll(line, "\n", "")
		args = strings.Split(line, ",")
		for i, s := range args {
			intArgs[i], _ = strconv.Atoi(s)
		}
		left, right := intArgs[:2], intArgs[2:]
		if left[0] >= right[0] && left[1] <= right[1] ||
			(right[0] >= left[0] && right[1] <= left[1]) {
			total++
		}
	}
	return total
}

func main() {
	fmt.Println("Solution to part one =>", part1())
}
