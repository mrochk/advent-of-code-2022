package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	next    *Cell
	content byte
}

func NewCell(content byte) *Cell {
	return &Cell{
		next:    nil,
		content: content,
	}
}

type Stack struct {
	top *Cell
}

func NewStack() *Stack {
	return &Stack{
		top: nil,
	}
}

func (s *Stack) Top(letter byte) {
	if s.top == nil {
		s.top = NewCell(letter)
	} else {
		newTop := NewCell(letter)
		newTop.next = s.top
		s.top = newTop
	}
}

func (s *Stack) Pop() {
	s.top = s.top.next
}

func (s *Stack) GetTop() byte {
	if s.top != nil {
		return s.top.content
	}
	return 0
}

func (s *Stack) Print() {
	str := ""
	temp := s.top
	for temp != nil {
		str += string(temp.content)
		temp = temp.next
	}
	fmt.Println(str)
}

func getStacks() []*Stack {
	var (
		input, _   = os.Open("input.txt")
		reader     = bufio.NewReader(input)
		stacks     = make([]*Stack, 0)
		tempStacks = make([]*Stack, 0)
		values     = make([]byte, 0)
	)
	for i := 0; i < 9; i++ {
		stacks = append(stacks, NewStack())
		tempStacks = append(tempStacks, NewStack())
	}
	for {
		line, _ := reader.ReadString('\n')
		if line == " 1   2   3   4   5   6   7   8   9 \n" {
			break
		}
		for i := 1; i <= 33; i += 4 {
			values = append(values, line[i])
		}
		for i, letter := range values {
			if letter != ' ' {
				tempStacks[i].Top(letter)
			}
		}

		values = nil
	}
	for i := range stacks {
		for tempStacks[i].top != nil {
			stacks[i].Top(tempStacks[i].GetTop())
			tempStacks[i].Pop()
		}
	}
	return stacks
}

func part1() string {
	var (
		stacks   = getStacks()
		input, _ = os.Open("input.txt")
		reader   = bufio.NewReader(input)
	)
	for {
		line, _ := reader.ReadString('\n')
		if line == "\n" {
			break
		}
	}
	for {
		line, _ := reader.ReadString('\n')
		if line == "" {
			break
		}
		line = strings.ReplaceAll(line, "\n", "")
		args := strings.Split(line, " ")
		move, _ := strconv.Atoi(args[1])
		from, _ := strconv.Atoi(args[3])
		to, _ := strconv.Atoi(args[5])
		from = from - 1
		to = to - 1
		for ; move > 0; move-- {
			stacks[to].Top(stacks[from].GetTop())
			stacks[from].Pop()
		}
	}
	res := ""
	for i := range stacks {
		res += string(stacks[i].GetTop())
	}
	return res
}

func main() {
	fmt.Println("Solution to part one =>", part1())
}
