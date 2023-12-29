package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Module struct {
	mtype   byte
	name    string
	targets []string
}

func (m Module) String() string {
	switch m.mtype {
	case 0:
		return fmt.Sprintf("%s -> %v", m.name, m.targets)
	case '%':
		return fmt.Sprintf("%%%s: %v -> %v", m.name, ffStates[m.name], m.targets)
	case '&':
		return fmt.Sprintf("&%s: %v -> %v", m.name, nandStates[m.name], m.targets)
	case 'd':
		return fmt.Sprintf("%s: %v -> %v", m.name, dummyStates[m.name], m.targets)
	}
	return ""
}

type Pulse struct {
	source string
	target string
	signal bool
}

func (p Pulse) String() string {
	var signal string
	if p.signal {
		signal = "high"
	} else {
		signal = "low"
	}
	return fmt.Sprintf("%s -%s-> %s", p.source, signal, p.target)
}

var circuit = make(map[string]Module)
var reversed = make(map[string][]string)
var ffStates = make(map[string]bool)
var nandStates = make(map[string]map[string]bool)
var dummyStates = make(map[string]bool)

func readInput(f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		tokens := strings.Split(line, " -> ")
		m := Module{}
		m.name = tokens[0]
		if m.name[0] == '%' || m.name[0] == '&' {
			m.mtype = m.name[0]
			m.name = m.name[1:]
		}
		m.targets = strings.Split(tokens[1], ", ")
		circuit[m.name] = m
		for _, t := range m.targets {
			if _, ok := circuit[t]; !ok {
				circuit[t] = Module{mtype: 'd', name: t}
			}
			sources, ok := reversed[t]
			if !ok {
				sources = []string{m.name}
			} else {
				sources = append(sources, m.name)
			}
			reversed[t] = sources
		}
	}
}

func initStates() {
	for _, m := range circuit {
		if m.mtype == '%' {
			ffStates[m.name] = false
		}
		if m.mtype == 'd' {
			dummyStates[m.name] = false
		}
		for _, t := range m.targets {
			tm, ok := circuit[t]
			if ok && tm.mtype == '&' {
				state := nandStates[tm.name]
				if state == nil {
					state = make(map[string]bool)
				}
				state[m.name] = false
				nandStates[tm.name] = state
			}
		}
	}
}

func pushButton(queue []Pulse) (int, int) {
	numLows, numHighs := 0, 0
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		// fmt.Println(p)
		if p.signal {
			numHighs++
		} else {
			numLows++
		}
		m, ok := circuit[p.target]
		if !ok {
			continue
		}
		switch m.mtype {
		case 0:
			for _, t := range m.targets {
				queue = append(queue, Pulse{m.name, t, p.signal})
			}
		case '%':
			if !p.signal {
				state := !ffStates[m.name]
				for _, t := range m.targets {
					queue = append(queue, Pulse{m.name, t, state})
				}
				ffStates[m.name] = state
			}
		case '&':
			nandStates[m.name][p.source] = p.signal
			signal := true
			for _, s := range nandStates[m.name] {
				signal = signal && s
			}
			for _, t := range m.targets {
				queue = append(queue, Pulse{m.name, t, !signal})
			}
		case 'd':
			dummyStates[m.name] = p.signal
			// fmt.Println("dummy", m)
		}
	}
	return numLows, numHighs
}

func firstPart(f *os.File) {
	readInput(f)
	fmt.Println(circuit)
	initStates()
	numLows, numHighs := 0, 0
	for i := 0; i < 1000; i++ {
		pulseQueue := []Pulse{Pulse{"button", "broadcaster", false}}
		l, h := pushButton(pulseQueue)
		fmt.Println(ffStates)
		fmt.Println(nandStates)
		numLows += l
		numHighs += h
		fmt.Println(numLows, numHighs)
	}
	fmt.Println(numLows * numHighs)
}

func gcd(a, b int) int {
	if a == b {
		return a
	}
	mx, mn := max(a, b), min(a, b)
	return gcd(mn, mx-mn)
}

func lcm(nums map[string]int) int {
	ans := 1
	for _, v := range nums {
		ans = (ans * v) / gcd(ans, v)
	}
	return ans
}

func secondPart(f *os.File) {
	readInput(f)
	fmt.Println(circuit)
	fmt.Println(reversed)
	initStates()
	fmt.Println()
	queue := []string{"rx"}
	for i := 0; i < 2 && len(queue) != 0; i++ {
		p := queue[0]
		queue = queue[1:]
		for _, t := range reversed[p] {
			m := circuit[t]
			fmt.Println(m)
		}
		fmt.Println()
		queue = append(queue, reversed[p]...)
	}
	numSources := len(reversed["df"])
	prevCounts := make(map[string]int)
	cycleCounts := make(map[string]int)
outer:
	for i := 1; i < 1000000000; i++ {
		pulseQueue := []Pulse{Pulse{"button", "broadcaster", false}}
		for len(pulseQueue) != 0 {
			p := pulseQueue[0]
			pulseQueue = pulseQueue[1:]
			// fmt.Println(p)
			m, ok := circuit[p.target]
			if !ok {
				continue
			}
			switch m.mtype {
			case 0:
				for _, t := range m.targets {
					pulseQueue = append(pulseQueue, Pulse{m.name, t, p.signal})
				}
			case '%':
				if !p.signal {
					state := !ffStates[m.name]
					for _, t := range m.targets {
						pulseQueue = append(pulseQueue, Pulse{m.name, t, state})
					}
					ffStates[m.name] = state
				}
			case '&':
				nandStates[m.name][p.source] = p.signal
				signal := true
				for _, s := range nandStates[m.name] {
					signal = signal && s
				}
				for _, t := range m.targets {
					pulseQueue = append(pulseQueue, Pulse{m.name, t, !signal})
				}
				if p.target == "df" && p.signal {
					cycleCounts[p.source] = i - prevCounts[p.source]
					fmt.Println(p.source, cycleCounts[p.source])
					prevCounts[p.source] = i
					if len(cycleCounts) == numSources {
						break outer
					}
				}
			case 'd':
				dummyStates[m.name] = p.signal
				// fmt.Println("dummy", m)
			}
		}
	}
	fmt.Println(lcm(cycleCounts))
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
