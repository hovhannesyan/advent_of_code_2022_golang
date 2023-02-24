package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkWin(a, b uint8) int {
	b -= 23
	switch {
	case a == b:
		return 3
	case a == 65 && b == 66:
		return 6
	case a == 66 && b == 67:
		return 6
	case a == 67 && b == 65:
		return 6
	default:
		return 0
	}
}

func checkWinByInstructions(a uint8, b *uint8) int {
	switch *b {
	case 89:
		*b = a
	case 90:
		if a != 67 {
			*b = a + 1
		} else {
			*b = a - 2
		}
	case 88:
		if a != 65 {
			*b = a - 1
		} else {
			*b = a + 2
		}
	}

	*b += 23
	return checkWin(a, *b)
}

func checkChoice(b uint8) int {
	switch {
	case b == 89:
		return 2
	case b == 90:
		return 3
	default:
		return 1
	}
}

func main() {
	var (
		scoreByInstructions, score int
		tmp                        uint8
	)
	readFile, _ := os.Open("./day_2/input.txt")
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		score += checkWin(fileScanner.Text()[0], fileScanner.Text()[2]) + checkChoice(fileScanner.Text()[2])
		tmp = fileScanner.Text()[2]
		scoreByInstructions += checkWinByInstructions(fileScanner.Text()[0], &tmp) + checkChoice(tmp)
	}
	fmt.Println("Score: ", score)
	fmt.Println("Score: ", scoreByInstructions)
}
