package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func firstPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ": ")
		if len(game) < 2 {
			continue
		}
		gameId, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
		fmt.Println(gameId)
		records := strings.Split(game[1], "; ")
		impossible := false
	outer:
		for _, i := range records {
			counts := strings.Split(i, ", ")
			for _, ii := range counts {
				colors := strings.Split(ii, " ")
				count, _ := strconv.Atoi(colors[0])
				color := colors[1]
				fmt.Println(color, count)
				if (color == "red" && count > 12) || (color == "green" && count > 13) || (color == "blue" && count > 14) {
					impossible = true
					break outer
				}
			}
		}
		if !impossible {
			sum += gameId
		}
	}
	fmt.Println(sum)
}

func secondPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ": ")
		if len(game) < 2 {
			continue
		}
		records := strings.Split(game[1], "; ")
		maxRed, maxGreen, maxBlue := 0, 0, 0
		for _, i := range records {
			counts := strings.Split(i, ", ")
			for _, ii := range counts {
				colors := strings.Split(ii, " ")
				count, _ := strconv.Atoi(colors[0])
				color := colors[1]
				fmt.Println(color, count)
				if color == "red" && count > maxRed {
					maxRed = count
				}
				if color == "green" && count > maxGreen {
					maxGreen = count
				}
				if color == "blue" && count > maxBlue {
					maxBlue = count
				}
			}
		}
		sum += maxRed * maxGreen * maxBlue
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
