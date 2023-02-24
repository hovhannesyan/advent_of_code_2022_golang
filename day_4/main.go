package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseString(s string) (int, int, int, int) {
	sepElves := strings.Split(s, ",")
	firstElf := strings.Split(sepElves[0], "-")
	secondElf := strings.Split(sepElves[1], "-")
	firstElfStart, _ := strconv.Atoi(firstElf[0])
	firstElfEnd, _ := strconv.Atoi(firstElf[1])
	secondElfStart, _ := strconv.Atoi(secondElf[0])
	secondElfEnd, _ := strconv.Atoi(secondElf[1])

	return firstElfStart, firstElfEnd, secondElfStart, secondElfEnd
}

func checkOverlap(firstElfStart, firstElfEnd, secondElfStart, secondElfEnd int) int {
	if secondElfStart >= firstElfStart && secondElfStart <= firstElfEnd {
		return 1
	}
	if secondElfEnd >= firstElfStart && secondElfEnd <= firstElfEnd {
		return 1
	}
	if firstElfStart >= secondElfStart && firstElfStart <= secondElfEnd {
		return 1
	}
	if firstElfEnd >= secondElfStart && firstElfEnd <= secondElfEnd {
		return 1
	}
	return 0
}

func checkFullOverlap(firstElfStart, firstElfEnd, secondElfStart, secondElfEnd int) int {
	if (firstElfStart <= secondElfStart && firstElfEnd >= secondElfEnd) || (firstElfStart >= secondElfStart && firstElfEnd <= secondElfEnd) {
		return 1
	}
	return 0
}

func main() {
	var countFullOverlaps, countOverlaps int

	readFile, _ := os.Open("./day_4/input.txt")
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		countFullOverlaps += checkFullOverlap(parseString(fileScanner.Text()))
		countOverlaps += checkOverlap(parseString(fileScanner.Text()))
	}
	fmt.Println(countFullOverlaps)
	fmt.Println(countOverlaps)
}
