package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type nTree struct {
	Name     string
	Memory   int
	Parent   *nTree
	Children []*nTree
}

func addNTree(parent *nTree, name string) {
	parent.Children = append(parent.Children, &nTree{Name: name, Parent: parent})
}

func addFileToNTree(parent *nTree, name string, memory int) {
	parent.Children = append(parent.Children, &nTree{Name: name, Parent: parent, Memory: memory})
}

func findNTree(parent *nTree, name string) *nTree {
	if parent == nil {
		return nil
	}
	for _, child := range parent.Children {
		if child.Name == name {
			return child
		}
	}
	return nil
}

func goToHead(parent **nTree) {
	if (*parent).Parent == nil {
		return
	}
	*parent = (*parent).Parent
	goToHead(*&parent)
}

func sumMemoryByDir(head *nTree) {
	if len(head.Children) == 0 {
		return
	}
	var memory int
	for _, child := range head.Children {
		sumMemoryByDir(child)
		if child.Memory > 0 {
			memory += child.Memory
		}
	}
	head.Memory = memory
}

func checkMemory(head *nTree, checked map[int]int) {
	if len(head.Children) == 0 {
		return
	}
	checked[head.Memory]++
	for _, child := range head.Children {
		if len(child.Children) > 0 {
			checkMemory(child, checked)
		}
	}
}

func createNTree(parent **nTree, s string) {
	strs := strings.Split(s, " ")

	if len(strs) == 3 {
		if strs[2] == ".." {
			*parent = (*parent).Parent
		} else {
			*parent = findNTree(*parent, strs[2])
		}
	} else {
		if strs[1] == "ls" {
			return
		}
		if strs[0] == "dir" {
			addNTree(*parent, strs[1])
		} else {
			tmp, _ := strconv.Atoi(strs[0])
			addFileToNTree(*parent, strs[1], tmp)
		}

	}
}

func main() {
	var (
		iterations int
		head       = &nTree{
			Name:   "/",
			Memory: 0,
		}
		checked    = make(map[int]int)
		checkedMin int
	)
	readFile, _ := os.Open("./day_7/input.txt")
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if iterations == 0 {
			iterations++
			continue
		}

		//fmt.Println(iterations)

		createNTree(&head, fileScanner.Text())

		iterations++
	}

	goToHead(&head)
	sumMemoryByDir(head)
	checkMemory(head, checked)
	checkedMin = head.Memory
	for val, _ := range checked {
		if val >= (head.Memory-40000000) && val < checkedMin {
			checkedMin = val
		}
	}
	fmt.Println(checkedMin)
}
