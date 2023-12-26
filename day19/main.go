package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	category  byte
	condition byte
	value     int
	dest      string
}

func (r Rule) String() string {
	if r.category == 0 {
		return fmt.Sprintf("%s", r.dest)
	}
	return fmt.Sprintf("%c%c%v:%s", r.category, r.condition, r.value, r.dest)
}

type Workflow struct {
	name  string
	rules []Rule
}

func (w Workflow) String() string {
	return fmt.Sprintf("%s: %v", w.name, w.rules)
}

type Part map[byte]int

func (p Part) String() string {
	return fmt.Sprintf("x=%v, m=%v, a=%v, s=%v", p['x'], p['m'], p['a'], p['s'])
}

var workflows = make(map[string]Workflow)
var parts = make([]Part, 0)

func readInput(f *os.File) {
	re := regexp.MustCompile("(.+)\\{(.+)\\}")
	re2 := regexp.MustCompile("(.)([><]+)(\\d+):(.*)")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		rules := make([]Rule, 0)
		matches := re.FindStringSubmatch(line)
		for _, rule := range strings.Split(matches[2], ",") {
			matches2 := re2.FindStringSubmatch(rule)
			if len(matches2) > 0 {
				value, _ := strconv.Atoi(matches2[3])
				rules = append(rules, Rule{matches2[1][0], matches2[2][0], value, matches2[4]})
			} else {
				rules = append(rules, Rule{dest: rule})
			}
		}
		workflows[matches[1]] = Workflow{matches[1], rules}
	}
	workflows["A"] = Workflow{name: "A"}
	workflows["R"] = Workflow{name: "R"}
	fmt.Println(workflows)
	re3 := regexp.MustCompile("\\{(.)=(\\d+),(.)=(\\d+),(.)=(\\d+),(.)=(\\d+)\\}")
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		matches3 := re3.FindStringSubmatch(line)
		part := Part{}
		for i := 1; i < len(matches3); i += 2 {
			value, _ := strconv.Atoi(matches3[i+1])
			part[matches3[i][0]] = value
		}
		parts = append(parts, part)
	}
	fmt.Println(parts)
}

func firstPart(f *os.File) {
	readInput(f)
	ans := 0
	for _, p := range parts {
		fmt.Println("part", p)
		w := workflows["in"]
		for {
			var dest string
			for _, r := range w.rules {
				fmt.Println("rule", r)
				v, ok := p[r.category]
				if ok {
					if r.condition == '>' && v > r.value {
						dest = r.dest
						break
					}
					if r.condition == '<' && v < r.value {
						dest = r.dest
						break
					}

				} else {
					dest = r.dest
					break
				}
			}
			fmt.Println(dest)
			if dest == "R" {
				fmt.Println(p, "rejected")
				break
			}
			if dest == "A" {
				fmt.Println(p, "accepted")
				ans += p['x'] + p['m'] + p['a'] + p['s']
				break
			}
			w = workflows[dest]
		}
	}
	fmt.Println(ans)
}

var acceptPaths = make([][]Rule, 0)

func dfs(w Workflow, path []Rule) {
	if w.name == "R" {
		return
	}
	if w.name == "A" {
		acceptPaths = append(acceptPaths, path)
		return
	}
	lastPath := make([]Rule, len(path))
	copy(lastPath, path)
	for _, r := range w.rules {
		fmt.Println("r", r)
		fmt.Println("path", path)
		if r.category == 0 {
			dfs(workflows[r.dest], lastPath)
		} else {
			copyPath := make([]Rule, len(lastPath)+1)
			copy(copyPath, lastPath)
			copyPath[len(copyPath)-1] = r
			dfs(workflows[r.dest], copyPath)
			var condition byte
			var value int
			switch r.condition {
			case '<':
				condition = '>'
				value = r.value - 1
			case '>':
				condition = '<'
				value = r.value + 1
			}
			lastPath = append(lastPath, Rule{r.category, condition, value, ""})
		}
	}
}

func secondPart(f *os.File) {
	readInput(f)
	dfs(workflows["in"], make([]Rule, 0))
	fmt.Println(len(acceptPaths))
	ans := 0
	for _, p := range acceptPaths {
		fmt.Println(p)
		minValues := map[byte]int{'x': 1, 'm': 1, 'a': 1, 's': 1}
		maxValues := map[byte]int{'x': 4001, 'm': 4001, 'a': 4001, 's': 4001}
		for _, r := range p {
			if r.condition == '>' {
				minValues[r.category] = max(r.value+1, minValues[r.category])
			}
			if r.condition == '<' {
				maxValues[r.category] = min(r.value, maxValues[r.category])
			}
		}
		fmt.Printf("%4v<=x<%4v, %4v<=m<%4v, %4v<=a<%4v, %4v<=s<%4v\n", minValues['x'], maxValues['x'], minValues['m'], maxValues['m'], minValues['a'], maxValues['a'], minValues['s'], maxValues['s'])
		if minValues['x'] > maxValues['x'] {
			continue
		}
		if minValues['m'] > maxValues['m'] {
			continue
		}
		if minValues['a'] > maxValues['a'] {
			continue
		}
		if minValues['s'] > maxValues['s'] {
			continue
		}
		ans += (maxValues['x'] - minValues['x']) * (maxValues['m'] - minValues['m']) * (maxValues['a'] - minValues['a']) * (maxValues['s'] - minValues['s'])
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
