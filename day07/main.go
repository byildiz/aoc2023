package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards    string
	strength int
	bid      int
}

var cardToHex = map[rune]rune{
	'A': 'E',
	'K': 'D',
	'Q': 'C',
	'J': 'B',
	'T': 'A',
	'9': '9',
	'8': '8',
	'7': '7',
	'6': '6',
	'5': '5',
	'4': '4',
	'3': '3',
	'2': '2',
}

var cardToHex2 = map[rune]rune{
	'A': 'E',
	'K': 'D',
	'Q': 'C',
	'J': '1',
	'T': 'A',
	'9': '9',
	'8': '8',
	'7': '7',
	'6': '6',
	'5': '5',
	'4': '4',
	'3': '3',
	'2': '2',
}

func getHandStrength(cards string) int64 {
	runeToCount := make(map[rune]int)
	for _, c := range cards {
		runeToCount[c]++
	}
	counts := make([]int, len(runeToCount))
	index := 0
	for _, v := range runeToCount {
		counts[index] = v
		index++
	}
	sort.Slice(counts, func(i, j int) bool { return counts[i] > counts[j] })
	strength, _ := strconv.ParseInt(cards, 16, 0)
	if counts[0] == 5 {
		fmt.Println("Five of a kind") // 6
		strength += 6 << 21
	} else if counts[0] == 4 {
		fmt.Println("Four of a kind") // 5
		strength += 5 << 21
	} else if counts[0] == 3 && counts[1] == 2 {
		fmt.Println("Full house") // 4
		strength += 4 << 21
	} else if counts[0] == 3 {
		fmt.Println("Three of a kind") // 3
		strength += 3 << 21
	} else if counts[0] == 2 && counts[1] == 2 {
		fmt.Println("Two pair") // 2
		strength += 2 << 21
	} else if counts[0] == 2 {
		fmt.Println("One pair") // 1
		strength += 1 << 21
	} else {
		fmt.Println("High card") // 0
	}
	return strength
}

func getHandStrength2(cards string) int64 {
	runeToCount := make(map[rune]int)
	for _, c := range cards {
		runeToCount[c]++
	}
	counts := make([]int, len(runeToCount))
	index := 0
	for k, v := range runeToCount {
		if k != 'J' {
			counts[index] = v
			index++
		}
	}
	sort.Slice(counts, func(i, j int) bool { return counts[i] > counts[j] })
	// in order to make the hand as stronger as possible, we should convert jokers to the most frequent card
	counts[0] += runeToCount['J']
	chars := make([]rune, len(cards))
	for i, c := range cards {
		chars[i] = cardToHex2[c]
	}
	strength, _ := strconv.ParseInt(string(chars), 16, 0)
	if counts[0] == 5 {
		fmt.Println("Five of a kind") // 6
		strength += 6 << 21
	} else if counts[0] == 4 {
		fmt.Println("Four of a kind") // 5
		strength += 5 << 21
	} else if counts[0] == 3 && counts[1] == 2 {
		fmt.Println("Full house") // 4
		strength += 4 << 21
	} else if counts[0] == 3 {
		fmt.Println("Three of a kind") // 3
		strength += 3 << 21
	} else if counts[0] == 2 && counts[1] == 2 {
		fmt.Println("Two pair") // 2
		strength += 2 << 21
	} else if counts[0] == 2 {
		fmt.Println("One pair") // 1
		strength += 1 << 21
	} else {
		fmt.Println("High card") // 0
	}
	return strength
}

func firstPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	hands := make([]Hand, 0)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		if len(tokens) == 1 {
			break
		}
		chars := make([]rune, len(tokens[0]))
		for i, c := range tokens[0] {
			chars[i] = cardToHex[c]
		}
		cards := string(chars)
		strength := int(getHandStrength(cards))
		bid, _ := strconv.Atoi(tokens[1])
		fmt.Println(cards, strength, bid)
		hands = append(hands, Hand{cards, strength, bid})
	}
	fmt.Println(hands)
	sort.SliceStable(hands, func(i, j int) bool { return hands[i].strength < hands[j].strength })
	fmt.Println(hands)
	answer := 0
	for i, h := range hands {
		answer += (i + 1) * h.bid
	}
	fmt.Println("answer", answer)
}

func secondPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	hands := make([]Hand, 0)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		if len(tokens) == 1 {
			break
		}
		strength := int(getHandStrength2(tokens[0]))
		bid, _ := strconv.Atoi(tokens[1])
		fmt.Println(tokens[0], strength, bid)
		hands = append(hands, Hand{tokens[0], strength, bid})
	}
	fmt.Println(hands)
	sort.SliceStable(hands, func(i, j int) bool { return hands[i].strength < hands[j].strength })
	fmt.Println(hands)
	answer := 0
	for i, h := range hands {
		answer += (i + 1) * h.bid
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
