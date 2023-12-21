package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type T [][]rune

func (t T) String() string {
	var b strings.Builder
	for _, col := range t {
		b.WriteString(string(col))
	}
	return b.String()
}

func readInput(f *os.File) T {
	scanner := bufio.NewScanner(f)
	rows := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		rows = append(rows, line)
	}
	// transpose the input to make things easy
	cols := make(T, len(rows[0]))
	for i := 0; i < len(rows[0]); i++ {
		cols[i] = make([]rune, len(rows))
		for j := 0; j < len(rows); j++ {
			cols[i][j] = rune(rows[j][i])
		}
	}
	return cols
}

func printTable(t T) {
	rows, cols := len(t), len(t[0])
	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			fmt.Printf("%c", t[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func calcLoad(t [][]rune) int {
	ans := 0
	for _, col := range t {
		row := len(col)
		load := 0
		for i, c := range col {
			if c == 'O' {
				load += row - i
			}
		}
		ans += load
	}
	return ans
}

func firstPart(f *os.File) {
	table := readInput(f)
	tiltTable(table, 'w')
	ans := calcLoad(table)
	fmt.Println(ans)
}

var dp map[string]int

func tiltTable(t T, d rune) {
	rows, cols := len(t), len(t[0])
	switch d {
	case 'w':
		for i := 0; i < rows; i++ {
			k := 0
			for j := 0; j < cols; j++ {
				c := t[i][j]
				if c == 'O' {
					t[i][k], t[i][j] = c, t[i][k]
					k++
				}
				if c == '#' {
					k = j + 1
				}
			}
		}
	case 'e':
		for i := 0; i < rows; i++ {
			k := cols - 1
			for j := cols - 1; j >= 0; j-- {
				c := t[i][j]
				if c == 'O' {
					t[i][k], t[i][j] = c, t[i][k]
					k--
				}
				if c == '#' {
					k = j - 1
				}
			}
		}
	case 'n':
		for j := 0; j < cols; j++ {
			k := 0
			for i := 0; i < rows; i++ {
				c := t[i][j]
				if c == 'O' {
					t[k][j], t[i][j] = c, t[k][j]
					k++
				}
				if c == '#' {
					k = i + 1
				}
			}
		}
	case 's':
		for j := 0; j < cols; j++ {
			k := rows - 1
			for i := rows - 1; i >= 0; i-- {
				c := t[i][j]
				if c == 'O' {
					t[k][j], t[i][j] = c, t[k][j]
					k--
				}
				if c == '#' {
					k = i - 1
				}
			}
		}
	}
}

func secondPart(f *os.File) {
	table := readInput(f)
	printTable(table)
	first, second, cycle := -1, -1, -1
	dp = make(map[string]int)
	i := 0
	for ; i < 1000000000; i++ {
		k := table.String()
		// fmt.Println(k)
		v, ok := dp[k]
		if ok {
			if first == -1 {
				first = v
			} else if second == -1 && first == v {
				second = i
				cycle = second - first
				fmt.Println(first, second, cycle)
				break
			}
		} else {
			dp[k] = i
		}
		for _, d := range [4]rune{'w', 'n', 'e', 's'} {
			tiltTable(table, d)
			// printTable(table)
		}
		// fmt.Println(k, v, calcLoad(table))
	}
	r := (1000000000 - i) % cycle
	for j := 0; j < r; j++ {
		for _, d := range [4]rune{'w', 'n', 'e', 's'} {
			tiltTable(table, d)
			// printTable(table)
		}
	}
	fmt.Println(calcLoad(table))
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
