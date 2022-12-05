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

func (s *Stack) Val() byte {
	if s.top != nil {
		return s.top.content
	}
	return 0
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
			stacks[i].Top(tempStacks[i].Val())
			tempStacks[i].Pop()
		}
	}
	return stacks
}

func solution() (string, string) {
	var (
		stacks1  = getStacks()
		stacks2  = getStacks()
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
		var (
			args      = strings.Split(line, " ")
			move, _   = strconv.Atoi(args[1])
			from, _   = strconv.Atoi(args[3])
			to, _     = strconv.Atoi(args[5])
			tempStack = NewStack()
		)
		from, to = from-1, to-1
		for i := move; i > 0; i-- {
			tempStack.Top(stacks2[from].Val())
			stacks2[from].Pop()
		}
		for ; move > 0; move-- {
			stacks1[to].Top(stacks1[from].Val())
			stacks1[from].Pop()
			stacks2[to].Top(tempStack.Val())
			tempStack.Pop()
		}
	}
	res1, res2 := "", ""
	for i := range stacks1 {
		res1 += string(stacks1[i].Val())
		res2 += string(stacks2[i].Val())
	}
	return res1, res2
}

func main() {
	p1, p2 := solution()
	fmt.Println("Solution to part one =>", p1)
	fmt.Println("Solution to part one =>", p2)
}
