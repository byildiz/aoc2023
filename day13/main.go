package main

import (
	"bufio"
	"fmt"
	"os"
)

func parsePatterns(patterns []string) [2][]int {
	rows := make([]int, len(patterns))
	for i, p := range patterns {
		fmt.Println(p)
		n := 0
		for _, c := range p {
			n *= 2
			if c == '#' {
				n += 1
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
		rows[i] = n
	}
	cols := make([]int, len(patterns[0]))
	for i := 0; i < len(patterns[0]); i++ {
		n := 0
		for j := 0; j < len(patterns); j++ {
			n *= 2
			if patterns[j][i] == '#' {
				n += 1
			}
		}
		cols[i] = n
	}
	return [2][]int{rows, cols}
}

func readInput(f *os.File) [][2][]int {
	scanner := bufio.NewScanner(f)
	inputs := make([][2][]int, 0)
	patterns := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			input := parsePatterns(patterns)
			inputs = append(inputs, input)
			patterns = make([]string, 0)
		} else {
			patterns = append(patterns, line)
		}
	}
	return inputs
}

func findMirror(nums []int) int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		found := true
		j, k := i, i+1
		for ; j >= 0 && k < l; j, k = j-1, k+1 {
			if nums[j] != nums[k] {
				found = false
				break
			}
		}
		if found {
			return i + 1
		}
	}
	return -1
}

func firstPart(f *os.File) {
	inputs := readInput(f)
	var total, mirror int
	for _, n := range inputs {
		fmt.Println(n[0])
		mirror = findMirror(n[0])
		fmt.Println(mirror)
		if mirror > 0 {
			total += mirror * 100
		}
		fmt.Println(n[1])
		mirror = findMirror(n[1])
		fmt.Println(mirror)
		if mirror > 0 {
			total += mirror
		}
	}
	fmt.Println(total)
}

func equal(d int) bool {
	return (d & (d - 1)) == 0
}

func findMirror2(nums []int) int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		found := true
		one := 0
		j, k := i, i+1
		for ; j >= 0 && k < l; j, k = j-1, k+1 {
			d := nums[j] ^ nums[k]
			if d != 0 && equal(d) {
				one++
			}
			if !equal(d) {
				found = false
				break
			}
		}
		if found && one == 1 {
			return i + 1
		}
	}
	return -1
}

func secondPart(f *os.File) {
	inputs := readInput(f)
	var total, mirror int
	for _, n := range inputs {
		fmt.Println(n[0])
		mirror = findMirror2(n[0])
		fmt.Println(mirror)
		if mirror > 0 {
			total += mirror * 100
		}
		fmt.Println(n[1])
		mirror = findMirror2(n[1])
		fmt.Println(mirror)
		if mirror > 0 {
			total += mirror
		}
	}
	fmt.Println(total)
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
