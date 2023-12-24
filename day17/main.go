package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(f *os.File) [][]int {
	scanner := bufio.NewScanner(f)
	mat := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		vec := make([]int, len(line))
		for i, c := range line {
			vec[i] = int(c) - '0'
		}
		mat = append(mat, vec)
	}
	return mat
}

type Key struct {
	r, c   int
	dir    rune
	length int
}

type State struct {
	key  Key
	loss int
	prev *State
}

func (s State) String() string {
	return fmt.Sprintf("r=%v, c=%v, dir=%c, len=%v, loss:%v", s.key.r, s.key.c, s.key.dir, s.key.length, s.loss)
}

func print(mat *[][]int, rows, cols int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%4d", (*mat)[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func printPath(mat *[][]int, rows, cols int, s *State) {
	states := make([][]*State, rows)
	for i := 0; i < rows; i++ {
		states[i] = make([]*State, cols)
	}
	for s != nil {
		states[s.key.r][s.key.c] = s
		s = s.prev
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if states[i][j] != nil {
				fmt.Printf("%c", states[i][j].key.dir)
			} else {
				fmt.Printf("%d", (*mat)[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func bfs(mat *[][]int) int {
	rows, cols := len(*mat), len((*mat)[0])
	queue := make([]State, 1)
	queue[0] = State{Key{0, 0, 'x', 0}, -(*mat)[0][0], nil}
	dp := make(map[Key]int)
	losses := make([][]int, rows)
	for i := 0; i < rows; i++ {
		losses[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			losses[i][j] = (rows + cols) * 9
		}
	}
	losses[0][0] = 0
	var minState *State
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		key, loss := s.key, s.loss
		r, c, dir, length := key.r, key.c, key.dir, key.length
		// fmt.Println(s)
		// fmt.Println(queue)
		// print(&losses, rows, cols)
		// fmt.Scanln()
		if r == rows || r < 0 || c == cols || c < 0 {
			continue
		}
		currLoss := loss + (*mat)[r][c]
		if currLoss < losses[r][c] {
			losses[r][c] = currLoss
			if r == rows-1 && c == cols-1 {
				minState = &s
			}
		}
		prevLoss, ok := dp[key]
		if !ok || currLoss < prevLoss {
			dp[key] = currLoss
			if dir != 'r' && (dir != 'l' || length < 3) {
				l := 1
				if dir == 'l' {
					l = length + 1
				}
				queue = append(queue, State{Key{r, c + 1, 'l', l}, currLoss, &s})
			}
			if dir != 'u' && (dir != 'd' || length < 3) {
				l := 1
				if dir == 'd' {
					l = length + 1
				}
				queue = append(queue, State{Key{r + 1, c, 'd', l}, currLoss, &s})
			}
			if dir != 'l' && (dir != 'r' || length < 3) {
				l := 1
				if dir == 'r' {
					l = length + 1
				}
				queue = append(queue, State{Key{r, c - 1, 'r', l}, currLoss, &s})
			}
			if dir != 'd' && (dir != 'u' || length < 3) {
				l := 1
				if dir == 'u' {
					l = length + 1
				}
				queue = append(queue, State{Key{r - 1, c, 'u', l}, currLoss, &s})
			}
		}
	}
	printPath(mat, rows, cols, minState)
	return losses[rows-1][cols-1]
}

func firstPart(f *os.File) {
	mat := readInput(f)
	fmt.Println(mat)
	fmt.Println(bfs(&mat))
}

func bfs2(mat *[][]int) int {
	rows, cols := len(*mat), len((*mat)[0])
	queue := make([]State, 2)
	queue[0] = State{Key{0, 1, 'l', 1}, 0, nil}
	queue[1] = State{Key{1, 0, 'd', 1}, 0, nil}
	dp := make(map[Key]int)
	losses := make([][]int, rows)
	for i := 0; i < rows; i++ {
		losses[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			losses[i][j] = (rows + cols) * 9
		}
	}
	losses[0][0] = 0
	var minState *State
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		key, loss := s.key, s.loss
		r, c, dir, length := key.r, key.c, key.dir, key.length
		// fmt.Println(s)
		// fmt.Println(queue)
		// print(&losses, rows, cols)
		// fmt.Scanln()
		if r == rows || r < 0 || c == cols || c < 0 || length > 10 {
			continue
		}
		currLoss := loss + (*mat)[r][c]
		if currLoss < losses[r][c] && length >= 4 {
			losses[r][c] = currLoss
			if r == rows-1 && c == cols-1 {
				minState = &s
			}
		}
		prevLoss, ok := dp[key]
		if !ok || currLoss < prevLoss {
			dp[key] = currLoss
			if dir != 'r' && (dir == 'l' || length >= 4) {
				l := length
				if dir != 'l' {
					l = 0
				}
				queue = append(queue, State{Key{r, c + 1, 'l', l + 1}, currLoss, &s})
			}
			if dir != 'u' && (dir == 'd' || length >= 4) {
				l := length
				if dir != 'd' {
					l = 0
				}
				queue = append(queue, State{Key{r + 1, c, 'd', l + 1}, currLoss, &s})
			}
			if dir != 'l' && (dir == 'r' || length >= 4) {
				l := length
				if dir != 'r' {
					l = 0
				}
				queue = append(queue, State{Key{r, c - 1, 'r', l + 1}, currLoss, &s})
			}
			if dir != 'd' && (dir == 'u' || length >= 4) {
				l := length
				if dir != 'u' {
					l = 0
				}
				queue = append(queue, State{Key{r - 1, c, 'u', l + 1}, currLoss, &s})
			}
		}
	}
	printPath(mat, rows, cols, minState)
	return losses[rows-1][cols-1]
}

func secondPart(f *os.File) {
	mat := readInput(f)
	fmt.Println(mat)
	fmt.Println(bfs2(&mat))
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
