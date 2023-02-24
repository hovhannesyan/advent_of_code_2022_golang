package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []rune

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str rune) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return ' ', false
	} else {
		element := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return element, true
	}
}

func parseInputStacks(s []string) []Stack {
	var (
		stacks []Stack
		i      int
		strI   = 1
		m      = make(map[int]int)
	)
	for _, val := range s[len(s)-1] {
		if val == ' ' {
			continue
		}
		stacks = append(stacks, Stack{})
		i, _ = strconv.Atoi(string(val))
		m[strI] = i - 1
		strI += 4
	}
	for i := len(s) - 2; i >= 0; i-- {
		for j, val := range s[i] {
			if stackNum, ok := m[j]; ok && val != ' ' {
				stacks[stackNum].Push(val)
			}
		}
	}
	return stacks
}

func getLastCrates(stacks []Stack) string {
	var lastCrates string
	for _, stack := range stacks {
		tmp, _ := stack.Pop()
		lastCrates += string(tmp)
		stack.Push(tmp)
	}
	return lastCrates
}

func parseInstruction(str string) (int, int, int) {
	var (
		count, from, to int
		ints            []string
	)
	str = strings.ReplaceAll(str, "move", "")
	str = strings.ReplaceAll(str, "from", "")
	str = strings.ReplaceAll(str, "to", "")
	ints = strings.Split(str, "  ")
	count, _ = strconv.Atoi(strings.ReplaceAll(ints[0], " ", ""))
	from, _ = strconv.Atoi(ints[1])
	to, _ = strconv.Atoi(ints[2])
	return count, from, to
}

func moveCrates(stack []Stack, count, from, to int) {
	for i := 0; i < count; i++ {
		tmp, _ := stack[from-1].Pop()
		stack[to-1].Push(tmp)
	}
}

func moveCratesWithNewCrane(stack []Stack, count, from, to int) {
	var tmp Stack
	for i := 0; i < count; i++ {
		tmpVal, _ := stack[from-1].Pop()
		tmp.Push(tmpVal)
	}
	for i := 0; i < count; i++ {
		tmpVal, _ := tmp.Pop()
		stack[to-1].Push(tmpVal)
	}
}

func main() {
	var (
		stacksInput []string
		stacks      []Stack
	)
	readFile, _ := os.Open("./day_5/input.txt")
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			continue
		}
		if len(stacksInput) < 9 {
			stacksInput = append(stacksInput, fileScanner.Text())
			continue
		} else if len(stacksInput) == 9 {
			stacks = parseInputStacks(stacksInput)
			stacksInput = append(stacksInput, "")
		}
		count, from, to := parseInstruction(fileScanner.Text())
		moveCratesWithNewCrane(stacks, count, from, to)
	}
	fmt.Println(getLastCrates(stacks))
}
