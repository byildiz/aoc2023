package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func findNumber(j int, line string) (int, int) {
	if !isDigit(line[j]) {
		return -1, -1
	}
	start := j
	for ; start >= 0 && isDigit(line[start]); start-- {
	}
	start += 1
	end := j
	for ; end < len(line) && isDigit(line[end]); end++ {
	}
	n, _ := strconv.Atoi(line[start:end])
	return start, n
}

func firstPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	data := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		data = append(data, line)
	}
	rows, cols := len(data), len(data[0])
	fmt.Println(rows, cols)
	numbers := make(map[int]int)
	for i, line := range data {
		for j, c := range line {
			// find symbols
			if c != '.' && (c < '0' || c > '9') {
				if i < rows-1 {
					s, n := findNumber(j, data[i+1])
					if n >= 0 {
						numbers[(i+1)*cols+s] = n
					}
				}
				if i > 0 {
					s, n := findNumber(j, data[i-1])
					if n >= 0 {
						numbers[(i-1)*cols+s] = n
					}
				}
				if j < cols-1 {
					s, n := findNumber(j+1, data[i])
					if n >= 0 {
						numbers[(i)*cols+s] = n
					}
				}
				if j > 0 {
					s, n := findNumber(j-1, data[i])
					if n >= 0 {
						numbers[(i)*cols+s] = n
					}
				}
				if i < rows-1 && j < cols-1 {
					s, n := findNumber(j+1, data[i+1])
					if n >= 0 {
						numbers[(i+1)*cols+s] = n
					}
				}
				if i > 0 && j > 0 {
					s, n := findNumber(j-1, data[i-1])
					if n >= 0 {
						numbers[(i-1)*cols+s] = n
					}
				}
				if i > 0 && j < cols-1 {
					s, n := findNumber(j+1, data[i-1])
					if n >= 0 {
						numbers[(i-1)*cols+s] = n
					}
				}
				if i < rows-1 && j > 0 {
					s, n := findNumber(j-1, data[i+1])
					if n >= 0 {
						numbers[(i+1)*cols+s] = n
					}
				}
			}
		}
	}
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Println(sum)
}

func secondPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	data := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		data = append(data, line)
	}
	rows, cols := len(data), len(data[0])
	fmt.Println(rows, cols)
	sum := 0
	for i, line := range data {
		for j, c := range line {
			// find gears
			if c == '*' {
				numbers := make(map[int]int)
				if i < rows-1 {
					s, n := findNumber(j, data[i+1])
					if n >= 0 {
						numbers[(i+1)*cols+s] = n
					}
				}
				if i > 0 {
					s, n := findNumber(j, data[i-1])
					if n >= 0 {
						numbers[(i-1)*cols+s] = n
					}
				}
				if j < cols-1 {
					s, n := findNumber(j+1, data[i])
					if n >= 0 {
						numbers[(i)*cols+s] = n
					}
				}
				if j > 0 {
					s, n := findNumber(j-1, data[i])
					if n >= 0 {
						numbers[(i)*cols+s] = n
					}
				}
				if i < rows-1 && j < cols-1 {
					s, n := findNumber(j+1, data[i+1])
					if n >= 0 {
						numbers[(i+1)*cols+s] = n
					}
				}
				if i > 0 && j > 0 {
					s, n := findNumber(j-1, data[i-1])
					if n >= 0 {
						numbers[(i-1)*cols+s] = n
					}
				}
				if i > 0 && j < cols-1 {
					s, n := findNumber(j+1, data[i-1])
					if n >= 0 {
						numbers[(i-1)*cols+s] = n
					}
				}
				if i < rows-1 && j > 0 {
					s, n := findNumber(j-1, data[i+1])
					if n >= 0 {
						numbers[(i+1)*cols+s] = n
					}
				}
				if len(numbers) == 2 {
					ratio := 1
					for _, n := range numbers {
						ratio *= n
					}
					sum += ratio
				}
			}
		}
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
