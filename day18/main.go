package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Line struct {
	dir    byte
	length int
	color  string
}

func readInput(f *os.File) []Line {
	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile("([RLUD]{1}) ([0-9]+) \\((.*)\\)")
	lines := make([]Line, 0)
	for scanner.Scan() {
		l := scanner.Text()
		matches := re.FindStringSubmatch(l)
		if len(matches) == 4 {
			length, _ := strconv.Atoi(matches[2])
			lines = append(lines, Line{matches[1][0], length, matches[3]})
		}
	}
	return lines
}

func correctInput(lines []Line) []Line {
	for i := range lines {
		l := &lines[i]
		switch l.color[6] {
		case '0':
			l.dir = 'R'
		case '1':
			l.dir = 'D'
		case '2':
			l.dir = 'L'
		case '3':
			l.dir = 'U'
		}
		length, _ := strconv.ParseInt(l.color[1:6], 16, 0)
		l.length = int(length)
	}
	return lines
}

func printMat(mat *[][]bool, rows, cols int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (*mat)[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func dfsArea(mat *[][]bool, rows, cols, x, y int) (int, bool) {
	if x < 0 || x == cols || y < 0 || y == rows {
		return 0, false
	}
	if (*mat)[y][x] {
		return 0, true
	}
	(*mat)[y][x] = true
	aR, inR := dfsArea(mat, rows, cols, x+1, y)
	aL, inL := dfsArea(mat, rows, cols, x-1, y)
	aD, inD := dfsArea(mat, rows, cols, x, y+1)
	aU, inU := dfsArea(mat, rows, cols, x, y-1)
	return 1 + aR + aL + aD + aU, inR && inL && inD && inU
}

func getPerimeter(lines *[]Line) int {
	perimeter := 0
	for _, l := range *lines {
		perimeter += l.length
	}
	return perimeter
}

func getPoints(lines []Line) [][2]int {
	n := len(lines)
	x, y := 0, 0
	ps := make([][2]int, n)
	for i, l := range lines {
		switch l.dir {
		case 'R':
			x += l.length
		case 'D':
			y += l.length
		case 'L':
			x -= l.length
		case 'U':
			y -= l.length
		}
		ps[i] = [2]int{x, y}
	}
	return ps
}

func firstPart(f *os.File) {
	lines := readInput(f)
	fmt.Println(lines)
	perimeter := getPerimeter(&lines)
	fmt.Println(perimeter)
	n := len(lines)
	ps := getPoints(lines)
	fmt.Println(ps)
	minX, minY, maxX, maxY := n, n, -n, -n
	for _, p := range ps {
		minX = min(minX, p[0])
		minY = min(minY, p[1])
		maxX = max(maxX, p[0])
		maxY = max(maxY, p[1])
	}
	rows := maxY - minY + 1
	cols := maxX - minX + 1
	mat := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		mat[i] = make([]bool, cols)
	}
	for i := 0; i < n; i++ {
		p1, p2 := ps[(i-1+n)%n], ps[i]
		x1, y1 := p1[0]-minX, p1[1]-minY
		x2, y2 := p2[0]-minX, p2[1]-minY
		if x1 == x2 {
			for j := min(y1, y2); j <= max(y1, y2); j++ {
				mat[j][x1] = true
			}
		}
		if y1 == y2 {
			for j := min(x1, x2); j <= max(x1, x2); j++ {
				mat[y1][j] = true
			}
		}
	}
	printMat(&mat, rows, cols)
	area, ok := 0, false
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !mat[i][j] {
				area, ok = dfsArea(&mat, rows, cols, j, i)
				printMat(&mat, rows, cols)
				if ok {
					break
				}
			}
		}
		if ok {
			break
		}
	}
	fmt.Println(perimeter + area)
}

func shoelaceArea(points *[][2]int) int {
	n := len(*points)
	area := 0
	for i := range *points {
		p1, p2 := (*points)[i], (*points)[(i+1)%n]
		area += p1[0]*p2[1] - p1[1]*p2[0]
	}
	return area / 2
}

func greenArea(lines *[]Line) int {
	area, y := 0, 0
	for _, l := range *lines {
		switch l.dir {
		case 'L':
			area -= l.length * y
		case 'R':
			area += l.length * y
		case 'U':
			y += l.length
		case 'D':
			y -= l.length
		}
	}
	return area
}

func secondPart(f *os.File) {
	lines := readInput(f)
	fmt.Println(lines)
	lines2 := correctInput(lines)
	fmt.Println(lines2)
	// n := len(lines)
	a := greenArea(&lines)
	ps := getPoints(lines)
	a2 := shoelaceArea(&ps)
	fmt.Println(a == a2)
	fmt.Println(a, a2)
	p := getPerimeter(&lines)
	fmt.Println(a + p/2 + 1)
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
