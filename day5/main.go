package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func firstPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	src := make([]int, 0)
	for _, n := range strings.Split(strings.Split(scanner.Text(), ": ")[1], " ") {
		nn, _ := strconv.Atoi(n)
		src = append(src, nn)
	}
	r, _ := regexp.Compile("(\\d+) (\\d+) (\\d+)")
	scanner.Scan()
	for scanner.Scan() {
		fmt.Println("src", src)
		srcToDst := make([][3]int, 0)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("line", line)
			matches := r.FindStringSubmatch(line)
			if len(matches) != 4 {
				break
			}
			numbers := [3]int{}
			for i, n := range matches[1:] {
				nn, _ := strconv.Atoi(n)
				numbers[i] = nn
			}
			srcToDst = append(srcToDst, numbers)
		}
		fmt.Println("map", srcToDst)
		dst := make([]int, len(src))
		for i, s := range src {
			d := -1
			for _, numbers := range srcToDst {
				if s >= numbers[1] && s < numbers[1]+numbers[2] {
					d = numbers[0] + s - numbers[1]
				}
			}
			if d == -1 {
				d = s
			}
			dst[i] = d
		}
		fmt.Println("dst", dst)
		src = dst
	}
	fmt.Println("answer", slices.Min(src))
}

func secondPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	src := make([][2]int, 0)
	pair := [2]int{}
	for i, n := range strings.Split(strings.Split(scanner.Text(), ": ")[1], " ") {
		nn, _ := strconv.Atoi(n)
		pair[i%2] = nn
		if i%2 == 1 {
			pair[1] += pair[0]
			src = append(src, pair)
		}
	}
	fmt.Println("src", src)
	r, _ := regexp.Compile("(\\d+) (\\d+) (\\d+)")
	scanner.Scan()
	for scanner.Scan() {
		srcToDst := make([][3]int, 0)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("line", line)
			matches := r.FindStringSubmatch(line)
			if len(matches) != 4 {
				break
			}
			numbers := [3]int{}
			for i, n := range matches[1:] {
				nn, _ := strconv.Atoi(n)
				numbers[i] = nn
			}
			srcToDst = append(srcToDst, numbers)
		}
		fmt.Println("map", srcToDst)
		dst := make([][2]int, 0)
		for _, pair := range src {
			startS := pair[0]
			endS := pair[1]
			// find intersections between one source entry and all of the current mapping
			sections := make([][4]int, 0)
			for _, numbers := range srcToDst {
				startM := numbers[1]
				endM := numbers[1] + numbers[2]
				startI := max(startS, startM)
				endI := min(endS, endM)
				if startI < endI {
					startD := numbers[0] + startI - startM
					endD := numbers[0] + endI - startM
					sections = append(sections, [4]int{startI, endI, startD, endD})
				}
			}
			// sort the intersections by start point. this approach assumes there is no intersection between the intersections.
			sort.Slice(sections, func(i, j int) bool { return sections[i][0] < sections[j][0] })
			fmt.Println("sections", sections)
			mark := startS
			for _, s := range sections {
				// uncovered source section
				if mark < s[0] {
					dst = append(dst, [2]int{mark, s[0]})
				}
				dst = append(dst, [2]int{s[2], s[3]})
				mark = s[1]
			}
			if len(sections) == 0 {
				dst = append(dst, [2]int{startS, endS})
			}
		}
		fmt.Println("dst", dst)
		src = dst
	}
	minV := math.MaxInt
	for _, pair := range src {
		minV = min(minV, pair[0])
	}
	fmt.Println("answer", minV)
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
