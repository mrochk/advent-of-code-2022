package main

import (
	"bufio"
	"fmt"
	"os"
)

func getCharScore() map[byte]int {
	var (
		charScore = make(map[byte]int)
		count     = 1
	)
	for char := 'a'; char <= 'z'; char, count = char+1, count+1 {
		charScore[byte(char)] = count
	}
	for char := 'A'; char <= 'Z'; char, count = char+1, count+1 {
		charScore[byte(char)] = count
	}
	return charScore
}

func part1() int {
	var (
		charScore = getCharScore()
		input, _  = os.Open("input.txt")
		reader    = bufio.NewReader(input)
		repeated  byte
		total     = 0
	)
	for {
	exit:
		line, _ := reader.ReadString('\n')
		if line == "" {
			return total
		}
		var (
			firstPart  = line[:len(line)/2]
			secondPart = line[len(line)/2:]
		)
		for i := range firstPart {
			for y := range secondPart {
				if firstPart[i] == secondPart[y] {
					repeated = firstPart[i]
					total += charScore[repeated]
					goto exit
				}
			}
		}
	}
}

func part2() int {
	var (
		charScore = getCharScore()
		input, _  = os.Open("input.txt")
		reader    = bufio.NewReader(input)
		total     = 0
		repeated  byte
	)
	for {
	exit:
		line1, _ := reader.ReadString('\n')
		line2, _ := reader.ReadString('\n')
		line3, _ := reader.ReadString('\n')
		if line3 == "" {
			return total
		}
		for i := range line1 {
			for y := range line2 {
				for j := range line3 {
					if line1[i] == line2[y] && line1[i] == line3[j] {
						repeated = line2[y]
						total += charScore[repeated]
						goto exit
					}
				}
			}
		}
	}
}

func main() {
	fmt.Println("Solution to part 1 =>", part1())
	fmt.Println("Solution to part 2 =>", part2())
}
