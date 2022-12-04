package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution() (int, int) {
	var (
		input, _   = os.Open("input.txt")
		reader     = bufio.NewReader(input)
		res1, res2 = 0, 0
		line       string
		args       []string
		intArgs    [4]int
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
			res1++
		}
		if (left[0] >= right[1] && left[0] <= right[1]) ||
			(left[0] <= right[1] && left[1] >= right[0]) {
			res2++
		}
	}
	return res1, res2
}

func main() {
	p1, p2 := solution()
	fmt.Println("Solution to part one =>", p1)
	fmt.Println("Solution to part two =>", p2)
}
