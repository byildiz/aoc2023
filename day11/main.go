package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Coord struct {
	R, C int
}

func readInput(f *os.File) ([]Coord, []int, []int) {
	scanner := bufio.NewScanner(f)
	coords := make([]Coord, 0)
	emptyRows := make([]int, 0)
	nonEmptyCols := make(map[int]bool)
	cols := 0
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		cols = len(line)
		rowEmpty := true
		for j, c := range line {
			if c == '#' {
				coords = append(coords, Coord{i, j})
				rowEmpty = false
				nonEmptyCols[j] = true
			}
		}
		if rowEmpty {
			emptyRows = append(emptyRows, i)
		}
	}
	emptyCols := make([]int, 0)
	for i := 0; i < cols; i++ {
		if !nonEmptyCols[i] {
			emptyCols = append(emptyCols, i)
		}
	}
	return coords, emptyRows, emptyCols
}

func expandRows(coords []Coord, emptyRows []int, e int) {
	// do expantion reverse since the rows effects the rows coming after them
	r := len(coords) - 1
	for i := len(emptyRows) - 1; i >= 0; i-- {
		for ; coords[r].R > emptyRows[i]; r-- {
		}
		for j := r + 1; j < len(coords); j++ {
			coords[j].R += e - 1
		}
	}
}

func expandCols(coords []Coord, emptyCols []int, e int) {
	// do expantion reverse since the cols effects the cols coming after them
	c := len(coords) - 1
	for i := len(emptyCols) - 1; i >= 0; i-- {
		for ; coords[c].C > emptyCols[i]; c-- {
		}
		for j := c + 1; j < len(coords); j++ {
			coords[j].C += e - 1
		}
	}
}

func sumPaths(coords []Coord) int {
	total := 0
	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			c1, c2 := coords[i], coords[j]
			if c1.R > c2.R {
				total += c1.R - c2.R
			} else {
				total += c2.R - c1.R
			}
			if c1.C > c2.C {
				total += c1.C - c2.C
			} else {
				total += c2.C - c1.C
			}
		}
	}
	return total
}

func firstPart(f *os.File) {
	coords, emptyRows, emptyCols := readInput(f)
	fmt.Println("coords", coords)
	fmt.Println("emptyRows", emptyRows)
	fmt.Println("emptyCols", emptyCols)
	// coords are sorted by rows by default
	expandRows(coords, emptyRows, 1)
	fmt.Println("coords", coords)
	sort.Slice(coords, func(i, j int) bool { return coords[i].C < coords[j].C })
	expandCols(coords, emptyCols, 1)
	fmt.Println("coords", coords)
	answer := sumPaths(coords)
	fmt.Println("answer", answer)
}

func secondPart(f *os.File) {
	coords, emptyRows, emptyCols := readInput(f)
	fmt.Println("coords", coords)
	fmt.Println("emptyRows", emptyRows)
	fmt.Println("emptyCols", emptyCols)
	// coords are sorted by rows by default
	expandRows(coords, emptyRows, 1000000)
	fmt.Println("coords", coords)
	sort.Slice(coords, func(i, j int) bool { return coords[i].C < coords[j].C })
	expandCols(coords, emptyCols, 1000000)
	fmt.Println("coords", coords)
	answer := sumPaths(coords)
	fmt.Println("answer", answer)
}

func main() {
	inputPath := os.Args[1]
	fmt.Println(inputPath)
	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if len(os.Args) == 2 {
		fmt.Println("First part")
		firstPart(f)
	} else {
		fmt.Println("Second part")
		secondPart(f)
	}
}
