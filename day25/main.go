package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var E = map[string]map[string]bool{}

func readInput(f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		tokens := strings.Split(line, ": ")
		if _, ok := E[tokens[0]]; !ok {
			E[tokens[0]] = map[string]bool{}
		}
		targets := strings.Split(tokens[1], " ")
		for _, t := range targets {
			if _, ok := E[t]; !ok {
				E[t] = map[string]bool{}
			}
			E[tokens[0]][t] = true
			E[t][tokens[0]] = true
		}
	}
}

func dfs(s string, seen map[string]bool) int {
	if seen[s] {
		return 0
	}
	seen[s] = true
	count := 1
	for k, v := range E[s] {
		if v {
			count += dfs(k, seen)
		}
	}
	return count
}

func firstPart(f *os.File) {
	readInput(f)
	delete(E["pbx"], "njx")
	delete(E["njx"], "pbx")
	delete(E["zvk"], "sxx")
	delete(E["sxx"], "zvk")
	delete(E["sss"], "pzr")
	delete(E["pzr"], "sss")
	a := dfs("pbx", map[string]bool{})
	b := dfs("njx", map[string]bool{})
	fmt.Println(a,b)
	fmt.Println(a*b)
}

func secondPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
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
