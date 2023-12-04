package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func firstPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	points := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		splits := strings.Split(line, ": ")
		splits = strings.Split(splits[1], " | ")
		numbers := strings.Split(splits[0], " ")
		wNumbers := make(map[int]struct{})
		for _, n := range numbers {
			num, err := strconv.Atoi(n)
			if err == nil {
				wNumbers[num] = struct{}{}
			}
		}
		numbers = strings.Split(splits[1], " ")
		matches := 0
		for _, n := range numbers {
			num, err := strconv.Atoi(n)
			if err == nil {
				_, isIn := wNumbers[num]
				if isIn {
					matches++
				}
			}
		}
		if matches > 0 {
			points += 1 << (matches - 1)
		}
	}
	fmt.Println(points)
}

func secondPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	r, _ := regexp.Compile("Card\\s+(\\d+)\\: (.+) \\| (.+)")
	rr, _ := regexp.Compile("\\s+")
	counts := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		splits := r.FindStringSubmatch(line)
		if len(splits) == 0 {
			continue
		}
		cardId, _ := strconv.Atoi(splits[1])
		// original card
		counts[cardId]++
		numbers := rr.Split(splits[2], -1)
		wNumbers := make(map[int]struct{})
		for _, n := range numbers {
			num, err := strconv.Atoi(n)
			if err == nil {
				wNumbers[num] = struct{}{}
			}
		}
		numbers = rr.Split(splits[3], -1)
		matches := 0
		for _, n := range numbers {
			num, err := strconv.Atoi(n)
			if err == nil {
				_, isIn := wNumbers[num]
				if isIn {
					matches++
				}
			}
		}
		for i := cardId + 1; i <= cardId+matches; i++ {
			counts[i] += counts[cardId]
		}
	}
	sum := 0
	for _, v := range counts {
		sum += v
	}
	fmt.Println(sum)
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
