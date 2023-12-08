package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func firstPart(f *os.File) {
	regex1, _ := regexp.Compile(".*:\\s+(.*)")
	regex2, _ := regexp.Compile("\\s+")
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	timeStr := regex1.FindStringSubmatch(scanner.Text())[1]
	time := make([]int, 0)
	for _, x := range regex2.Split(timeStr, -1) {
		xx, _ := strconv.Atoi(x)
		time = append(time, xx)
	}
	fmt.Println("time", time)
	scanner.Scan()
	distanceStr := regex1.FindStringSubmatch(scanner.Text())[1]
	distance := make([]int, 0)
	for _, x := range regex2.Split(distanceStr, -1) {
		xx, _ := strconv.Atoi(x)
		distance = append(distance, xx)
	}
	fmt.Println("distance", distance)
	mul := 1
	for i := 0; i < len(time); i++ {
		t := time[i]
		d := distance[i]
		options := 0
		for j := 1; j < t; j++ {
			if j*(t-j) > d {
				options++
			}
		}
		mul *= options
	}
	fmt.Println("answer", mul)
}

func secondPart(f *os.File) {
	regex1, _ := regexp.Compile(".*:\\s+(.*)")
	regex2, _ := regexp.Compile("\\s+")
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	timeStr := regex1.FindStringSubmatch(scanner.Text())[1]
	time, _ := strconv.Atoi(regex2.ReplaceAllString(timeStr, ""))
	fmt.Println("time", time)
	scanner.Scan()
	distanceStr := regex1.FindStringSubmatch(scanner.Text())[1]
	distance, _ := strconv.Atoi(regex2.ReplaceAllString(distanceStr, ""))
	fmt.Println("distance", distance)
	start, end := 1, time/2
	for start < end {
		fmt.Println("points", start, end)
		m := (start + end) / 2
		d := m * (time - m)
		if d > distance {
			end = m - 1
		} else {
			start = m + 1
		}
	}
	// -1 is for 0
	fmt.Println("answer", time-1-2*start)
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
