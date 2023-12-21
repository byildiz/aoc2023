package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		return strings.Split(line, ",")
	}
	return nil
}

func hash(s string) int {
	cur := 0
	for _, c := range s {
		cur = ((cur + int(c)) * 17) % 256
	}
	return cur
}

func firstPart(f *os.File) {
	seq := readInput(f)
	fmt.Println(seq)
	ans := 0
	for _, s := range seq {
		fmt.Printf("'%s'\n", s)
		h := hash(s)
		fmt.Println(h)
		ans += h
	}
	fmt.Println(ans)
}

type Lens struct {
	label string
	focal int
	index int
}

func secondPart(f *os.File) {
	seq := readInput(f)
	fmt.Println(seq)
	m := [256]map[string]Lens{}
	idx := [256]int{}
	for i := 0; i < 256; i++ {
		m[i] = make(map[string]Lens)
	}
	for _, s := range seq {
		fmt.Printf("'%s'\n", s)
		if strings.HasSuffix(s, "-") {
			fmt.Println("deletion")
			s = s[:len(s)-1]
			h := hash(s)
			fmt.Printf("'%s': %v\n", s, h)
			delete(m[h], s)
		} else {
			fmt.Println("insertion")
			tokens := strings.Split(s, "=")
			s = tokens[0]
			v, _ := strconv.Atoi(tokens[1])
			h := hash(s)
			fmt.Printf("'%s': %v, %v\n", s, h, v)
			lens, ok := m[h][s]
			if ok {
				lens.focal = v
				m[h][s] = lens
			} else {
				m[h][s] = Lens{s, v, idx[h]}
				idx[h]++
			}
		}
	}
	// fmt.Println(m)
	ans := 0
	for i := 0; i < 256; i++ {
		if len(m[i]) > 0 {
			fmt.Println("Box", i, m[i])
		}
		lenses := make([]Lens, len(m[i]))
		j := 0
		for _, v := range m[i] {
			lenses[j] = v
			j++
		}
		sort.Slice(lenses, func(i, j int) bool { return lenses[i].index < lenses[j].index })
		power := 0
		for j = 0; j < len(lenses); j++ {
			power += (i + 1) * (j + 1) * lenses[j].focal
		}
		// fmt.Println(power)
		ans += power
	}
	fmt.Println(ans)
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
