package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func firstPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	ins := scanner.Text()
	fmt.Println(ins)
	re, _ := regexp.Compile("(\\w{3}) = \\((\\w{3}), (\\w{3})\\)")
	network := make(map[string]*Node)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) != 4 {
			continue
		}
		var node, left, right *Node
		if left = network[matches[2]]; left == nil {
			left = &Node{Name: matches[2]}
			network[matches[2]] = left
		}
		if right = network[matches[3]]; right == nil {
			right = &Node{Name: matches[3]}
			network[matches[3]] = right
		}
		if node = network[matches[1]]; node == nil {
			node = &Node{matches[1], left, right}
			network[matches[1]] = node
		} else {
			node.Left = left
			node.Right = right
		}
		fmt.Println(node)
	}
	steps := 0
	node := network["AAA"]
out:
	for {
		for _, i := range ins {
			if i == 'L' {
				node = node.Left
			} else {
				node = node.Right
			}
			steps++
			if node.Name == "ZZZ" {
				break out
			}
		}
	}
	fmt.Println("answer", steps)
}

func gcd(a, b int) int {
	for a != b {
		a, b = max(a, b), min(a, b)
		a -= b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func secondPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	ins := scanner.Text()
	fmt.Println(ins)
	re, _ := regexp.Compile("(\\w{3}) = \\((\\w{3}), (\\w{3})\\)")
	network := make(map[string]*Node)
	endsWithA := make([]*Node, 0)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) != 4 {
			continue
		}
		var node, left, right *Node
		if left = network[matches[2]]; left == nil {
			left = &Node{Name: matches[2]}
			network[matches[2]] = left
		}
		if right = network[matches[3]]; right == nil {
			right = &Node{Name: matches[3]}
			network[matches[3]] = right
		}
		if node = network[matches[1]]; node == nil {
			node = &Node{matches[1], left, right}
			network[matches[1]] = node
		} else {
			node.Left = left
			node.Right = right
		}
		if node.Name[2] == 'A' {
			endsWithA = append(endsWithA, node)
		}
		fmt.Println(node)
	}
	fmt.Println()
	firstSteps := make([]int, len(endsWithA))
	cycleSteps := make([]int, len(endsWithA))
	for j, a := range endsWithA {
		fmt.Println("a", a)
		steps := 0
		node := a
		var z *Node
	out:
		for {
			for _, i := range ins {
				if i == 'L' {
					node = node.Left
				} else {
					node = node.Right
				}
				steps++
				if node.Name[2] == 'Z' {
					fmt.Println("z", node)
					if node == z {
						cycleSteps[j] = steps
						break out
					}
					z = node
					firstSteps[j] = steps
					steps = 0
				}
			}
		}
	}
	fmt.Println("firstSteps", firstSteps)
	fmt.Println("cycleSteps", cycleSteps)
	answer := cycleSteps[0]
	for _, c := range cycleSteps {
		answer = lcm(answer, c)
	}
	fmt.Println("answer", answer)
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
