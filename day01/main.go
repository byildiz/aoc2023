package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func firstPart(f *os.File) {
	reader := bufio.NewReader(f)
	sum := 0
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if isPrefix {
			panic("Line is too long")
		}
		if len(line) <= 0 {
			continue
		}
		var first, last int
		var isFirst = true
		for _, c := range line {
			if c >= '0' && c <= '9' {
				if isFirst {
					first = int(c)
					isFirst = false
				}
				last = int(c)
			}
		}
		fmt.Printf("%c%c\n", first, last)
		sum += (first-'0')*10 + (last - '0')
	}
	fmt.Println(sum)
}

func secondPart(f *os.File) {
	digits := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	reader := bufio.NewReader(f)
	sum := 0
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if isPrefix {
			panic("Line is too long")
		}
		if len(line) <= 0 {
			continue
		}
		var first, last int
		var isFirst = true
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c >= '0' && c <= '9' {
				if isFirst {
					first = int(c - '0')
					isFirst = false
				}
				last = int(c - '0')
			} else {
				for j, d := range digits {
					if d == string(line[i:min(len(line), i+len(d))]) {
						if isFirst {
							first = j
							isFirst = false
						}
						last = j
					}
				}
			}
		}
		sum += first*10 + last
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
