package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		pointer, i int
		m          = make(map[string]int)
	)
	readFile, _ := os.Open("./day_6/input.txt")
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanRunes)

	for fileScanner.Scan() {
		if pointer != 14 {
			if _, ok := m[fileScanner.Text()]; !ok {
				m[fileScanner.Text()]++
				pointer++
			} else {
				pointer = 0
				for k := range m {
					delete(m, k)
				}
			}
		} else {
			fmt.Println(i)
			break
		}
		i++
	}

}
