package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	tiles := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		tiles = append(tiles, line)
	}
	return tiles
}

type S struct {
	x, y int
	dir  rune
}

func print(e *[][]bool, r, c int, s S) {
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == s.y && j == s.x {
				fmt.Print(string(s.dir))
			} else if (*e)[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

var dp map[S]bool

func dfs(tiles *[]string, energized *[][]bool, rows, cols int, s S) {
	x, y, dir := s.x, s.y, s.dir
	if x < 0 || x >= cols || y < 0 || y >= rows {
		return
	}
	if dp[s] {
		return
	} else {
		dp[s] = true
	}
	(*energized)[y][x] = true
	// print(energized, rows, cols, s)
	// fmt.Scanln()
	t := (*tiles)[y][x]
	up, left, down, right := false, false, false, false
	switch t {
	case '.':
		up = dir == 'u'
		left = dir == 'l'
		down = dir == 'd'
		right = dir == 'r'
	case '-':
		if dir == 'u' || dir == 'd' {
			left = true
			right = true
		} else {
			left = dir == 'l'
			right = dir == 'r'
		}
	case '|':
		if dir == 'l' || dir == 'r' {
			up = true
			down = true
		} else {
			up = dir == 'u'
			down = dir == 'd'
		}
	case '/':
		up = dir == 'r'
		left = dir == 'd'
		down = dir == 'l'
		right = dir == 'u'
	case '\\':
		up = dir == 'l'
		left = dir == 'u'
		down = dir == 'r'
		right = dir == 'd'
	}
	if up {
		dfs(tiles, energized, rows, cols, S{x, y - 1, 'u'})
	}
	if down {
		dfs(tiles, energized, rows, cols, S{x, y + 1, 'd'})
	}
	if left {
		dfs(tiles, energized, rows, cols, S{x - 1, y, 'l'})
	}
	if right {
		dfs(tiles, energized, rows, cols, S{x + 1, y, 'r'})
	}
}

func solve(tiles *[]string, rows, cols int, s S) int {
	energized := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		energized[i] = make([]bool, cols)
	}
	dp = make(map[S]bool)
	dfs(tiles, &energized, rows, cols, s)
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if energized[i][j] {
				ans++
			}
		}
	}
	return ans
}

func firstPart(f *os.File) {
	tiles := readInput(f)
	// fmt.Println(tiles)
	rows, cols := len(tiles), len(tiles[0])
	ans := solve(&tiles, rows, cols, S{0, 0, 'r'})
	fmt.Println(ans)
}

func secondPart(f *os.File) {
	tiles := readInput(f)
	// fmt.Println(tiles)
	rows, cols := len(tiles), len(tiles[0])
	ans := 0
	for i := 0; i < rows; i++ {
		ans = max(ans, solve(&tiles, rows, cols, S{0, i, 'r'}))
	}
	for i := 0; i < rows; i++ {
		ans = max(ans, solve(&tiles, rows, cols, S{cols - 1, i, 'l'}))
	}
	for i := 0; i < cols; i++ {
		ans = max(ans, solve(&tiles, rows, cols, S{i, 0, 'd'}))
	}
	for i := 0; i < cols; i++ {
		ans = max(ans, solve(&tiles, rows, cols, S{i, rows - 1, 'u'}))
	}
	fmt.Println(ans)
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
