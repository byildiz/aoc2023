package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type B struct {
	p1 [3]int
	p2 [3]int
}

var blocks []B

func readInput(f *os.File) {
	scanner := bufio.NewScanner(f)
	blocks = make([]B, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		points := strings.Split(line, "~")
		block := B{}
		p1 := strings.Split(points[0], ",")
		for i, v := range p1 {
			vv, _ := strconv.Atoi(v)
			block.p1[i] = vv
		}
		p2 := strings.Split(points[1], ",")
		for i, v := range p2 {
			vv, _ := strconv.Atoi(v)
			block.p2[i] = vv
		}
		blocks = append(blocks, block)
	}
}

var space [][][]int

// space indexing: z, y, x
func makeSpace() {
	maxs := [3]int{}
	for _, b := range blocks {
		for i := 0; i < 3; i++ {
			maxs[i] = max(b.p1[i]+1, b.p2[i]+1, maxs[i])
		}
	}
	fmt.Println(maxs)
	x, y, z := maxs[0], maxs[1], maxs[2]
	space = make([][][]int, z)
	for i := 0; i < z; i++ {
		space[i] = make([][]int, y)
		for j := 0; j < y; j++ {
			space[i][j] = make([]int, x)
		}
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			space[0][j][i] = -1
		}
	}
}

func drawBlock(b B, id int) {
	for z := b.p1[2]; z <= b.p2[2]; z++ {
		for y := b.p1[1]; y <= b.p2[1]; y++ {
			for x := b.p1[0]; x <= b.p2[0]; x++ {
				space[z][y][x] = id
			}
		}
	}
}

func deleteBlock(b B) {
	for z := b.p1[2]; z <= b.p2[2]; z++ {
		for y := b.p1[1]; y <= b.p2[1]; y++ {
			for x := b.p1[0]; x <= b.p2[0]; x++ {
				space[z][y][x] = 0
			}
		}
	}
}

func fillSpace() {
	for i, b := range blocks {
		drawBlock(b, i+1)
	}
}

func relocateBlock(id, diffZ int) {
	b := blocks[id-1]
	deleteBlock(b)
	b.p1[2] -= diffZ
	b.p2[2] -= diffZ
	blocks[id-1] = b
	drawBlock(b, id)
}

func fallBlock(id int) {
	b := blocks[id-1]
	minZ := min(b.p1[2], b.p2[2])
	z := minZ
outer:
	for ; ; z-- {
		for y := b.p1[1]; y <= b.p2[1]; y++ {
			for x := b.p1[0]; x <= b.p2[0]; x++ {
				v := space[z][y][x]
				// fmt.Println(v, id)
				if v != 0 && v != id {
					break outer
				}
			}
		}
	}
	diffZ := minZ - z - 1
	// fmt.Println(diffZ, "diff")
	if diffZ > 0 {
		relocateBlock(id, diffZ)
	}
}

func fallBlocks() {
	seen := make([]bool, len(blocks))
	for z := range space {
		for y := range space[z] {
			for x := range space[z][y] {
				v := space[z][y][x]
				if v > 0 && !seen[v-1] {
					// fmt.Println(v)
					fallBlock(v)
					// printSpace()
					seen[v-1] = true
				}
			}
		}
	}
}

func findSupporters() ([][]bool, []int) {
	supporters := make([][]bool, len(blocks))
	for i, b := range blocks {
		supporters[i] = make([]bool, len(blocks))
		maxZ := max(b.p1[2], b.p2[2]) + 1
		if maxZ >= len(space) {
			continue
		}
		for y := b.p1[1]; y <= b.p2[1]; y++ {
			for x := b.p1[0]; x <= b.p2[0]; x++ {
				v := space[maxZ][y][x]
				if v != 0 {
					// fmt.Println(i, v-1)
					supporters[i][v-1] = true
				}
			}
		}
	}
	numSupporters := make([]int, len(blocks))
	for i := range supporters {
		for j := range supporters[i] {
			if supporters[i][j] {
				numSupporters[j]++
			}
		}
	}
	return supporters, numSupporters
}

func printSpace() {
	for z := range space {
		for y := range space[z] {
			fmt.Println(space[z][y])
		}
		fmt.Println()
	}
	fmt.Println()
}

func firstPart(f *os.File) {
	readInput(f)
	fmt.Println(blocks)
	makeSpace()
	fillSpace()
	// printSpace()
	fallBlocks()
	// printSpace()
	fmt.Println(blocks)
	supporters, numSupporters := findSupporters()
	// fmt.Println(supporters)
	fmt.Println(numSupporters)
	ans := 0
	for i := range supporters {
		deintegratable := true
		for j := range supporters[i] {
			if supporters[i][j] && numSupporters[j] < 2 {
				// fmt.Println("no", i, j, numSupporters[j])
				deintegratable = false
				break
			}
		}
		if deintegratable {
			// fmt.Println("block", i+1)
			ans += 1
		}
	}
	fmt.Println(ans)
}

func bfs(i int, s [][]bool, ns []int) int {
	ans := 0
	queue := []int{i}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		for j, ss := range s[c] {
			if ss {
				ns[j]--
				if ns[j] == 0 {
					ans++
					queue = append(queue, j)
				}
			}
		}
	}
	return ans
}

func secondPart(f *os.File) {
	readInput(f)
	fmt.Println(blocks)
	makeSpace()
	fillSpace()
	// printSpace()
	fallBlocks()
	// printSpace()
	fmt.Println(blocks)
	supporters, numSupporters := findSupporters()
	// fmt.Println(supporters)
	fmt.Println(numSupporters)
	ans := 0
	seen := make([]bool, len(blocks))
	for z := range space {
		for y := range space[z] {
			for x := range space[z][y] {
				v := space[z][y][x]
				if v > 0 && !seen[v-1] {
					copyNumSupporters := make([]int, len(numSupporters))
					copy(copyNumSupporters, numSupporters)
					ans += bfs(v-1, supporters, copyNumSupporters)
					seen[v-1] = true
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
