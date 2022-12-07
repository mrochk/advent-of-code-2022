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
		dirsToSize = make(map[string]int)
		paths      = make([]string, 0)
		resP1      = 0
		resP2      = totalSpace
	)
	for {
		line, _ := reader.ReadString('\n')
		if line == "" {
			break
		}
		line = strings.Trim(line, "\n")
		args := strings.Split(line, " ")
		if args[0] == "$" {
			if args[1] == "cd" {
				if args[2] == ".." {
					paths = paths[:len(paths)-1]
				} else {
					paths = append(paths, args[2])
				}
			}
		} else {
			if args[0] != "dir" {
				size, _ := strconv.Atoi(args[0])
				path := ""
				for _, p := range paths {
					path += p
					dirsToSize[path] += size
				}
			}
		}
	}
	for _, size := range dirsToSize {
		if size <= atMostDirSize {
			resP1 += size
		}
	}
	toDelete := spaceRequired - (totalSpace - dirsToSize["/"])
	for _, dirSize := range dirsToSize {
		if dirSize >= toDelete && dirSize <= resP2 {
			resP2 = dirSize
		}
	}
	return resP1, resP2
}

const (
	atMostDirSize = 100_000
	totalSpace    = int(70e6)
	spaceRequired = int(30e6)
)

func main() {
	p1, p2 := solution()
	fmt.Println("Solution to part one =>", p1)
	fmt.Println("Solution to part two =>", p2)
}
