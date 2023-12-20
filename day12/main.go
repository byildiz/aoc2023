package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type S struct {
	row  []rune
	nums []int
	r, n int
	need []int
}

func (s S) String() string {
	return fmt.Sprintf("%s | %v | %v, %v | %v", string(s.row), s.nums, s.r, s.n, s.need)
}

func readInput(f *os.File) []S {
	scanner := bufio.NewScanner(f)
	states := make([]S, 0)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		if len(tokens) != 2 {
			break
		}
		row := []rune(tokens[0])
		nums := make([]int, 0)
		for _, c := range strings.Split(tokens[1], ",") {
			n, _ := strconv.Atoi(c)
			nums = append(nums, n)
		}
		states = append(states, S{row: row, nums: nums})
	}
	return states
}

// / this function assumes there is always a '?' in s.row
func canValid(s *S) bool {
	i := s.r
	for ; s.n < len(s.nums); s.n++ {
		n := s.nums[s.n]
		// skip all '.'
		for ; s.row[i] == '.'; i++ {
		}
		s.r = i
		// count '#'
		c := 0
		for ; s.row[i] == '#'; i++ {
			c++
		}
		// number of '#' should be n
		if s.row[i] == '.' && c != n {
			return false
		}
		if s.row[i] == '?' && c > n {
			return false
		}
		if s.row[i] == '?' {
			// check if there is enough space for rest of '#'
			if s.n < len(s.nums)-1 && len(s.row)-i < s.need[s.n+1] {
				fmt.Println(s, "no space")
				return false
			}
			return true
		}
		s.r = i
		// fmt.Println("i", i)
		// fmt.Printf("%c\n", s.row[i])
	}
	for ; i < len(s.row) && s.row[i] == '.'; i++ {
		s.r = i
	}
	for ; i < len(s.row) && (s.row[i] == '.' || s.row[i] == '?'); i++ {
	}
	if i == len(s.row) {
		return true
	}
	return false
}

func isValid(s S) bool {
	i := s.r
	for ; s.n < len(s.nums); s.n++ {
		n := s.nums[s.n]
		// skip all '.'
		for ; i < len(s.row) && s.row[i] == '.'; i++ {
		}
		// count '#'
		c := 0
		for ; i < len(s.row) && s.row[i] == '#'; i++ {
			c++
		}
		// number of '#' should be n
		if c != n {
			return false
		}
		// there should be at least one "." between groups
		if i < len(s.row) && s.row[i] != '.' {
			return false
		}
	}
	// the remaining should all be '.'
	for ; i < len(s.row); i++ {
		if s.row[i] != '.' {
			return false
		}
	}
	return true
}

func dfs(s S) int {
	i := 0
	for ; i < len(s.row) && s.row[i] != '?'; i++ {
	}
	// there is at lest one '?'
	if i < len(s.row) {
		// prune the rest if this state has already not valid
		if !canValid(&s) {
			// fmt.Println(s, "can't")
			// panic("end")
			return 0
		}
		var row []rune
		row = make([]rune, len(s.row))
		copy(row, s.row)
		row[i] = '#'
		r := dfs(S{row: row, nums: s.nums, r: s.r, n: s.n, need: s.need})
		row = make([]rune, len(s.row))
		copy(row, s.row)
		row[i] = '.'
		l := dfs(S{row: row, nums: s.nums, r: s.r, n: s.n, need: s.need})
		return l + r
	}
	if isValid(s) {
		// fmt.Println(s, "accepted")
		return 1
	}
	// fmt.Println(s, "rejected")
	return 0
}

func generateNeed(s S) []int {
	l := len(s.nums)
	need := make([]int, l)
	need[l-1] = s.nums[l-1]
	for i := l - 2; i >= 0; i-- {
		need[i] = need[i+1] + s.nums[i] + 1
	}
	return need
}

func firstPart(f *os.File) {
	states := readInput(f)
	fmt.Println("states", states)
	answer := 0
	for _, s := range states {
		s.need = generateNeed(s)
		fmt.Println(s)
		p := dfs(s)
		if p == 0 {
			panic(s)
		}
		answer += p
	}
	fmt.Println("answer", answer)
}

type S2 struct {
	ti, bi, c int
}

func (s S2) String() string {
	return fmt.Sprintf("ti=%v, bi=%v, c=%v", s.ti, s.bi, s.c)
}

var options = [2]rune{'.', '#'}
var dp map[S2]int

func dfs2(t []rune, b []int, s S2) int {
	// fmt.Println(s)
	total, exists := dp[s]
	if exists {
		return total
	}
	tl := len(t)
	bl := len(b)
	if s.ti == tl {
		if (s.bi == bl && s.c == 0) || (s.bi == bl-1 && b[s.bi] == s.c) {
			// fmt.Println(s, "accepted")
			return 1
		} else {
			// fmt.Println(s, "rejected")
			return 0
		}
	}
	for _, c := range options {
		// p := make([]rune, tl)
		// copy(p, s.p)
		var nextS S2
		if c == '#' && (t[s.ti] == '?' || t[s.ti] == '#') {
			// fmt.Println("#", string(t[s.ti]))
			nextS = S2{s.ti + 1, s.bi, s.c + 1}
			total += dfs2(t, b, nextS)
		} else if c == '.' && (t[s.ti] == '?' || t[s.ti] == '.') {
			// fmt.Println(".", string(t[s.ti]))
			if s.c > 0 && s.bi < bl && s.c == b[s.bi] {
				nextS = S2{s.ti + 1, s.bi + 1, 0}
				total += dfs2(t, b, nextS)
			} else if s.c == 0 {
				nextS = S2{s.ti + 1, s.bi, s.c}
				total += dfs2(t, b, nextS)
			}
		}
	}
	dp[s] = total
	return total
}

func secondPart(f *os.File) {
	states := readInput(f)
	fmt.Println("states", states)
	answer := 0
	for _, s := range states {
		// poor man's duplication logic
		row := make([]rune, 5*len(s.row)+4)
		nums := make([]int, 5*len(s.nums))
		for i := 0; i < 5; i++ {
			start := i*len(s.row) + i
			end := (i+1)*len(s.row) + i
			copy(row[start:end], s.row)
			if i != 4 {
				row[end] = '?'
			}
			copy(nums[i*len(s.nums):(i+1)*len(s.nums)], s.nums)
		}
		s.row = row
		s.nums = nums
		// s.need = generateNeed(s)
		fmt.Println(string(s.row))
		dp = make(map[S2]int)
		total := dfs2(s.row, s.nums, S2{0, 0, 0})
		fmt.Println(total)
		answer += total
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
