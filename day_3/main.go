package main

import (
	"bufio"
	"fmt"
	"os"
)

func findPrior(s string) int {
	firstCompartment := make(map[rune]int)

	for _, val := range s[:len(s)/2] {
		firstCompartment[val]++
	}
	for _, val := range s[len(s)/2:] {
		if _, ok := firstCompartment[val]; ok {
			return int(val)
		}
	}
	return 0
}

func findPriorUpdated(s []string) int {
	checkHash := make(map[rune]int)

	for _, val := range s[0] {
		if _, ok := checkHash[val]; !ok {
			checkHash[val]++
		}
	}
	for _, val := range s[1] {
		if _, ok := checkHash[val]; ok {
			checkHash[val]++
		}
	}
	for _, val := range s[2] {
		if i, ok := checkHash[val]; ok && i != 1 {
			return int(val)
		}
	}
	return 0
}

func getPriority(r int) int {
	if r >= 97 && r <= 122 {
		return r - 96
	}
	return r - 38
}

func main() {
	var (
		sum        int
		sumUpdated int
		slStr      []string
	)

	readFile, _ := os.Open("./day_3/input.txt")
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		prior := findPrior(fileScanner.Text())
		sum += getPriority(prior)
		slStr = append(slStr, fileScanner.Text())
		if len(slStr) == 3 {
			priorUpdated := findPriorUpdated(slStr)
			sumUpdated += getPriority(priorUpdated)
			slStr = slStr[:0]
		}
	}
	fmt.Println(sum)
	fmt.Println(sumUpdated)
}
