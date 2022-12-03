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
		line, _   = reader.ReadString('\n')
		repeated  byte
		total     = 0
	)
	for line != "" {
		var (
			firstPart  = line[:len(line)/2]
			secondPart = line[len(line)/2:]
		)
		for i := range firstPart {
			for y := range secondPart {
				if firstPart[i] == secondPart[y] {
					repeated = firstPart[i]
					break
				}
			}
		}
		total += charScore[repeated]
		line, _ = reader.ReadString('\n')
	}
	return total
}

func part2() int {
	var (
		charScore = getCharScore()
		input, _  = os.Open("input.txt")
		reader    = bufio.NewReader(input)
		line1, _  = reader.ReadString('\n')
		line2, _  = reader.ReadString('\n')
		line3, _  = reader.ReadString('\n')
		repeated  byte
		total     = 0
	)
	for line3 != "" {
		for i := range line1 {
			for y := range line2 {
				for j := range line3 {
					if line1[i] == line2[y] && line1[i] == line3[j] {
						repeated = line2[y]
						goto end
					}
				}
			}
		}
	end:
		total += charScore[repeated]
		line1, _ = reader.ReadString('\n')
		line2, _ = reader.ReadString('\n')
		line3, _ = reader.ReadString('\n')
	}
	return total
}

func main() {
	fmt.Println("Solution to part 1 =>", part1())
	fmt.Println("Solution to part 2 =>", part2())
}
