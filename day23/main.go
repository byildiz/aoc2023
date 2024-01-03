package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lines = []string{}
var rows, cols int

func readInput(f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}
	rows = len(lines)
	cols = len(lines[0])
}

func dfs(pr, pc, r, c int, seen map[[2]int]bool) (int, bool) {
	if r < 0 || r == rows || c < 0 || c == cols {
		return 0, false
	}
	v := lines[r][c]
	if v == '#' {
		return 0, false
	}
	key := [2]int{r, c}
	if seen[key] {
		return 0, false
	}
	if pr == r {
		if pc > c && v == '>' {
			return 0, false
		}
		if pc < c && v == '<' {
			return 0, false
		}
	}
	if pc == c {
		if pr > r && v == 'v' {
			return 0, false
		}
		if pr < r && v == '^' {
			return 0, false
		}
	}
	seen[key] = true
	if r == rows-1 {
		return 1, true
	}
	// fmt.Println(r, c)
	maxLen := 0
	for _, offset := range [4][2]int{[2]int{1, 0}, [2]int{-1, 0}, [2]int{0, 1}, [2]int{0, -1}} {
		copySeen := make(map[[2]int]bool, len(seen))
		for kk, vv := range seen {
			copySeen[kk] = vv
		}
		if l, ok := dfs(r, c, r+offset[0], c+offset[1], copySeen); ok {
			maxLen = max(maxLen, l)
		}
	}
	return maxLen + 1, true
}

func firstPart(f *os.File) {
	readInput(f)
	fmt.Println(lines)
	sc := strings.IndexByte(lines[0], '.')
	fmt.Println(dfs(0, sc, 1, sc, map[[2]int]bool{[2]int{0, sc}: true}))
}

type S struct {
	r, c, d, id int
}

var cache = map[[2]int]int{}

var V = map[[2]int]int{}
var E = map[int]map[int]int{}

func isPath(r, c int) bool {
	if r < 0 || r == rows || c < 0 || c == cols {
		return false
	}
	if lines[r][c] == '#' {
		return false
	}
	return true
}

func findWays(r, c int) [][2]int {
	ways := [][2]int{}
	if !isPath(r, c) {
		return ways
	}
	for _, offset := range [4][2]int{[2]int{1, 0}, [2]int{-1, 0}, [2]int{0, 1}, [2]int{0, -1}} {
		rr, cc := r+offset[0], c+offset[1]
		key := [2]int{rr, cc}
		if isPath(rr, cc) {
			ways = append(ways, key)
		}
	}
	return ways
}

func reduceGraph(sc, ec int) {
	id := 0
	seen := map[[2]int]bool{}
	queue := []S{S{0, sc, 0, id}}
	for len(queue) > 0 {
		// s := queue[0]
		// queue = queue[1:]
		s := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		r, c, d, i := s.r, s.c, s.d, s.id
		if !isPath(r, c) {
			continue
		}
		key := [2]int{r, c}
		seen[key] = true
		ways := findWays(r, c)
		if r == 0 && c == sc {
			fmt.Println("source vertex")
			V[key] = id
			E[id] = map[int]int{}
			id++
		} else if r == rows-1 && c == ec {
			fmt.Println("target vertex")
			V[key] = id
			E[id] = map[int]int{}
			E[id][i] = d + 1
			E[i][id] = d + 1
			d = 0
			i = id
			id++
		} else if len(ways) > 2 {
			fmt.Println("branch point")
			V[key] = id
			E[id] = map[int]int{}
			E[id][i] = d + 1
			E[i][id] = d + 1
			d = 0
			i = id
			id++
		}
		for _, w := range ways {
			if j, ok := V[w]; ok && d > 2 {
				fmt.Println("visited vertex")
				E[i][j] = d + 2
				E[j][i] = d + 2
			}
			if !seen[w] {
				queue = append(queue, S{w[0], w[1], d + 1, i})
			}
		}
		// printMap(seen)
		// fmt.Println("s", s)
		// fmt.Println("q", queue)
		// fmt.Println("ways", ways)
		// fmt.Println("V", V)
		// fmt.Println("E", E)
		// fmt.Scanln()
	}
}

func dfs2(v, t, d int, seen map[int]bool) int {
	if seen[v] {
		return -d
	}
	seen[v] = true
	if v == t {
		// fmt.Println("d", d)
		// fmt.Println("seen", seen)
		// fmt.Println("num", len(seen))
		return d - len(seen) + 1
	}
	maxDepth := 0
	for u, w := range E[v] {
		// fmt.Println(v, u, w)
		copySeen := map[int]bool{}
		for kk, vv := range seen {
			copySeen[kk] = vv
		}
		maxDepth = max(maxDepth, dfs2(u, t, w+d, copySeen))
	}
	return maxDepth
}

func printMap(seen map[[2]int]bool) {
	for i, line := range lines {
		for j, c := range line {
			if _, ok := seen[[2]int{i, j}]; ok {
				fmt.Print("O")
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func secondPart(f *os.File) {
	readInput(f)
	fmt.Println(lines)
	sc := strings.IndexByte(lines[0], '.')
	ec := strings.IndexByte(lines[rows-1], '.')
	reduceGraph(sc, ec)
	fmt.Println(V)
	fmt.Println(E)
	fmt.Println(dfs2(0, V[[2]int{rows - 1, ec}], 0, map[int]bool{}))
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
