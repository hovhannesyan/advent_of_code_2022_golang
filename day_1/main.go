package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		tmp    int
		sum    int
		maxSum int
	)
	bestElves := make(map[int]int)
	readFile, _ := os.Open("./day_1/input.txt")
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for i := 0; i < 3; i++ {
		_, _ = readFile.Seek(0, 0)
		fileScanner = bufio.NewScanner(readFile)
		for fileScanner.Scan() {
			if fileScanner.Text() == "" {
				if _, ok := bestElves[sum]; !ok && maxSum < sum {
					maxSum = sum
				}
				sum = 0
				continue
			}
			tmp, _ = strconv.Atoi(fileScanner.Text())
			sum += tmp
		}
		bestElves[maxSum] = i
		maxSum = 0
		sum = 0
	}
	maxSum = 0
	for sum, _ := range bestElves {
		maxSum += sum
	}
	fmt.Println(bestElves)
	fmt.Println(maxSum)

}
