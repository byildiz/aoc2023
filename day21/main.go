package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}
	return lines
}

func findS(lines []string) (int, int) {
	for r, line := range lines {
		if c := strings.IndexByte(line, 'S'); c >= 0 {
			return r, c
		}
	}
	return -1, -1
}

type S struct {
	r, c, d int
}

func firstPart(f *os.File) {
	lines := readInput(f)
	fmt.Println(lines)
	rows, cols := len(lines), len(lines[0])
	r, c := findS(lines)
	s := S{r, c, 0}
	states := make(map[S]bool)
	states[s] = true
	queue := []S{s}
	for len(queue) > 0 {
		// fmt.Println(queue)
		// fmt.Println(states)
		// s = queue[len(queue)-1]
		// queue = queue[:len(queue)-1]
		s = queue[0]
		queue = queue[1:]
		// fmt.Println("curr", s)
		r, c, d := s.r, s.c, s.d
		if r < 0 || r == rows || c < 0 || c == cols || lines[r][c] == '#' {
			delete(states, s)
			continue
		}
		if d == 64 {
			continue
		}
		delete(states, s)
		var nextS S
		nextS = S{r + 1, c, d + 1}
		if !states[nextS] {
			queue = append(queue, nextS)
			states[nextS] = true
		}
		nextS = S{r - 1, c, d + 1}
		if !states[nextS] {
			queue = append(queue, nextS)
			states[nextS] = true
		}
		nextS = S{r, c + 1, d + 1}
		if !states[nextS] {
			queue = append(queue, nextS)
			states[nextS] = true
		}
		nextS = S{r, c - 1, d + 1}
		if !states[nextS] {
			queue = append(queue, nextS)
			states[nextS] = true
		}
	}
	fmt.Println(len(states))
}

func printDist(dist [][]int) {
	for i := 0; i < len(dist); i++ {
		for j := 0; j < len(dist[i]); j++ {
			d := dist[i][j]
			if d == -1 {
				fmt.Printf("%4c", '#')
			} else {
				fmt.Printf("%4d", dist[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type T struct {
	r, c int
}

func makeTile(rows, cols int) [][]int {
	tile := make([][]int, rows)
	for i := 0; i < rows; i++ {
		tile[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			tile[i][j] = -1
		}
	}
	return tile
}

func bfs(lines []string, r, c, trows, tcols int) map[T][][]int {
	rows, cols := len(lines), len(lines[0])
	tiles := make(map[T][][]int)
	queue := [][5]int{[5]int{0, 0, r, c, 0}}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		// fmt.Println(s)
		tr, tc, r, c, d := s[0], s[1], s[2], s[3], s[4]
		if r < 0 {
			tr--
			r += rows
		}
		if r == rows {
			tr++
			r = 0
		}
		if c < 0 {
			tc--
			c += cols
		}
		if c == cols {
			tc++
			c = 0
		}
		if tr < -trows || tr > trows || tc < -tcols || tc > tcols {
			continue
		}
		if lines[r][c] == '#' {
			continue
		}
		k := T{tr, tc}
		tile, ok := tiles[k]
		if !ok {
			tile = makeTile(rows, cols)
			tiles[k] = tile
		}
		if tile[r][c] != -1 {
			continue
		}
		tile[r][c] = d
		queue = append(queue, [5]int{tr, tc, r + 1, c, d + 1})
		queue = append(queue, [5]int{tr, tc, r - 1, c, d + 1})
		queue = append(queue, [5]int{tr, tc, r, c + 1, d + 1})
		queue = append(queue, [5]int{tr, tc, r, c - 1, d + 1})
	}
	return tiles
}

type K struct {
	d int
	c bool
}

var dp = make(map[K]int)

func countRest(maxDist, d, step int, c bool) int {
	key := K{d, c}
	count, ok := dp[key]
	if ok {
		return count
	}
	count = 0
	for ii := 0; d+step*ii <= maxDist; ii++ {
		dd := d + step*ii
		if dd <= maxDist && dd%2 == maxDist%2 {
			if c {
				count += ii + 1
			} else {
				count++
			}
		}
	}
	dp[key] = count
	return count
}

func secondPart(f *os.File) {
	lines := readInput(f)
	fmt.Println(lines)
	rows, cols := len(lines), len(lines[0])
	fmt.Println(rows, cols)
	r, c := findS(lines)
	numTiles := 3
	maxDist := 26501365
	tiles := bfs(lines, r, c, numTiles, numTiles)
	middleTile := tiles[T{0, 0}]
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if middleTile[i][j] == -1 {
				continue
			}
			for tr := -numTiles; tr <= numTiles; tr++ {
				for tc := -numTiles; tc <= numTiles; tc++ {
					k := T{tr, tc}
					t := tiles[k]
					d := t[i][j]
					if (tr == -numTiles || tr == numTiles) && (tc == -numTiles || tc == numTiles) {
						ans += countRest(maxDist, d, rows, true)
					} else if (tr == -numTiles || tr == numTiles) || (tc == -numTiles || tc == numTiles) {
						ans += countRest(maxDist, d, rows, false)
					} else if d <= maxDist && d%2 == maxDist%2 {
						ans++
					}
				}
			}
		}
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
