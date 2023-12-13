package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findConnectedPipes(tiles [][]rune, rows, cols, x, y int) [][2]int {
	prev := tiles[y][x]
	pipes := make([][2]int, 0)
	if x > 0 && (prev == 'S' || prev == '-' || prev == 'J' || prev == '7') {
		next := tiles[y][x-1]
		if next == '-' || next == 'F' || next == 'L' || next == 'S' {
			// fmt.Printf("left: %c\n", next)
			pipes = append(pipes, [2]int{x - 1, y})
		}
	}
	if y > 0 && (prev == 'S' || prev == '|' || prev == 'J' || prev == 'L') {
		next := tiles[y-1][x]
		if next == '|' || next == 'F' || next == '7' || next == 'S' {
			// fmt.Printf("top: %c\n", next)
			pipes = append(pipes, [2]int{x, y - 1})
		}
	}
	if x < cols-1 && (prev == 'S' || prev == '-' || prev == 'F' || prev == 'L') {
		next := tiles[y][x+1]
		if next == '-' || next == '7' || next == 'J' || next == 'S' {
			// fmt.Printf("right: %c\n", next)
			pipes = append(pipes, [2]int{x + 1, y})
		}
	}
	if y < rows-1 && (prev == 'S' || prev == '|' || prev == 'F' || prev == '7') {
		next := tiles[y+1][x]
		if next == '|' || next == 'L' || next == 'J' || next == 'S' {
			// fmt.Printf("bottom: %c\n", next)
			pipes = append(pipes, [2]int{x, y + 1})
		}
	}
	return pipes
}

func printLoop(tiles [][]rune, loop [][2]int, length int) {
	for i := 0; i < length; i++ {
		x, y := loop[i][0], loop[i][1]
		fmt.Printf("%c", tiles[y][x])
	}
	fmt.Println()
}

func printTiles(tiles [][]rune) {
	for _, l := range tiles {
		fmt.Println(string(l))
	}
	fmt.Println()
}

func printOccupy(occupy [][]rune) {
	for _, l := range occupy {
		for _, c := range l {
			fmt.Printf("%v", c)
		}
		fmt.Println()
	}
	fmt.Println()
}

func expand(tiles [][]rune, loop [][2]int, length int) [][]rune {
	newTiles := make([][]rune, len(tiles)*2)
	for i := range newTiles {
		newTiles[i] = make([]rune, len(tiles[0])*2)
	}
	line := make([]rune, len(tiles[0])*2)
	dots := make([]rune, len(tiles[0])*2)
	for i := range dots {
		dots[i] = '.'
	}
	for i, l := range tiles {
		for j, c := range l {
			line[j*2] = c
			line[j*2+1] = '.'
		}
		copy(newTiles[i*2], line)
		copy(newTiles[i*2+1], dots)
	}
	for i := range loop[:length-1] {
		p, n := loop[i], loop[(i+1)%len(loop)]
		if p[0] == n[0] {
			y := min(p[1], n[1])*2 + 1
			newTiles[y][p[0]*2] = '|'
		} else {
			x := min(p[0], n[0])*2 + 1
			newTiles[p[1]*2][x] = '-'
		}
	}
	return newTiles
}

func dfs(tiles [][]rune, rows, cols, x, y int, depth int, loop [][2]int) int {
	loop[depth] = [2]int{x, y}
	// printLoop(tiles, loop, depth+1)
	if tiles[y][x] == 'S' {
		return depth + 1
	}
	for _, n := range findConnectedPipes(tiles, rows, cols, x, y) {
		// don't go back
		if depth > 0 && n == loop[depth-1] {
			continue
		}
		length := dfs(tiles, rows, cols, n[0], n[1], depth+1, loop)
		if length > -1 {
			return length
		}
	}
	return -1
}

func readInput(f *os.File) [][]rune {
	scanner := bufio.NewScanner(f)
	tiles := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		tiles = append(tiles, []rune(line))
	}
	return tiles
}

func findStart(tiles [][]rune) (int, int) {
	var x, y int
	for y = 0; y < len(tiles); y++ {
		x = strings.IndexRune(string(tiles[y]), 'S')
		if x != -1 {
			break
		}
	}
	return x, y
}

func fill(occupy [][]rune, rows, cols, x, y int, f rune) {
	if occupy[y][x] != 0 {
		return
	}
	occupy[y][x] = f
	if x > 0 {
		fill(occupy, rows, cols, x-1, y, f)
	}
	if y > 0 {
		fill(occupy, rows, cols, x, y-1, f)
	}
	if x < cols-1 {
		fill(occupy, rows, cols, x+1, y, f)
	}
	if y < rows-1 {
		fill(occupy, rows, cols, x, y+1, f)
	}
}

func firstPart(f *os.File) {
	tiles := readInput(f)
	x, y := findStart(tiles)
	// fmt.Println("start", x, y)
	rows, cols := len(tiles), len(tiles[0])
	loop := make([][2]int, rows*cols)
	loop[0] = [2]int{x, y}
	pipes := findConnectedPipes(tiles, rows, cols, x, y)
	// fmt.Println("S pipes", pipes)
	x, y = pipes[0][0], pipes[0][1]
	length := dfs(tiles, rows, cols, x, y, 1, loop)
	// fmt.Println("loop", loop[:length])
	fmt.Println("answer", length/2)
}

func secondPart(f *os.File) {
	tiles := readInput(f)
	// printTiles(tiles)
	x, y := findStart(tiles)
	// fmt.Println("start", x, y)
	rows, cols := len(tiles), len(tiles[0])
	loop := make([][2]int, rows*cols)
	loop[0] = [2]int{x, y}
	pipes := findConnectedPipes(tiles, rows, cols, x, y)
	// fmt.Println("S pipes", pipes)
	x, y = pipes[0][0], pipes[0][1]
	length := dfs(tiles, rows, cols, x, y, 1, loop)
	// fmt.Println("length", length)
	// fmt.Println("loop", loop[:length])
	newTiles := expand(tiles, loop, length)
	// printTiles(newTiles)
	x, y = findStart(newTiles)
	// fmt.Println("start", x, y)
	rows, cols = len(newTiles), len(newTiles[0])
	loop = make([][2]int, rows*cols)
	loop[0] = [2]int{x, y}
	pipes = findConnectedPipes(newTiles, rows, cols, x, y)
	// fmt.Println("S pipes", pipes)
	x, y = pipes[0][0], pipes[0][1]
	length = dfs(newTiles, rows, cols, x, y, 1, loop)
	// fmt.Println("length", length)
	// fmt.Println("loop", loop[:length])
	occupy := make([][]rune, rows)
	for i := range occupy {
		occupy[i] = make([]rune, cols)
	}
	for _, n := range loop[:length-1] {
		occupy[n[1]][n[0]] = 2
	}
	// printOccupy(occupy)
	for i := 0; i < rows; i++ {
		fill(occupy, rows, cols, 0, i, 1)
		fill(occupy, rows, cols, cols-1, i, 1)
	}
	for i := 0; i < cols; i++ {
		fill(occupy, rows, cols, i, 0, 1)
		fill(occupy, rows, cols, i, rows-1, 1)
	}
	// printOccupy(occupy)
	// printTiles(tiles)
	answer := 0
	for i := 0; i < rows; i += 2 {
		for j := 0; j < cols; j += 2 {
			// fmt.Printf("%v", occupy[i][j])
			if occupy[i][j] == 0 {
				answer++
			}
		}
		// fmt.Println()
	}
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
