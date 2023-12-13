package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binomCoeffs(prevCoeffs []int) []int {
	numPrevCoeffs := len(prevCoeffs)
	coeffs := make([]int, numPrevCoeffs+1)
	coeffs[0] = 1
	if numPrevCoeffs == 0 {
		return coeffs
	}
	var i int
	for ; i < numPrevCoeffs-1; i++ {
		coeffs[i+1] = prevCoeffs[i] + prevCoeffs[i+1]
	}
	coeffs[i+1] = 1
	return coeffs
}

func solve(nums []int) int {
	coeffs := make([]int, 0)
	ans := 0
	for i := 0; i < len(nums); i++ {
		coeffs = binomCoeffs(coeffs)
		offset := len(nums) - 1
		sign := 1
		for j, c := range coeffs {
			ans += sign * c * nums[offset-j]
			sign *= -1
		}
	}
	return ans
}

func solve2(nums []int) int {
	coeffs := make([]int, 0)
	sign2 := 1
	ans := 0
	for i := 0; i < len(nums); i++ {
		coeffs = binomCoeffs(coeffs)
		offset := len(coeffs) - 1
		sign := 1
		diff := 0
		for j, c := range coeffs {
			diff += sign * c * nums[offset-j]
			sign *= -1
		}
		ans += sign2 * diff
		sign2 *= -1
	}
	return ans
}

func firstPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	answer := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		nums := make([]int, 0)
		for _, t := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(t)
			nums = append(nums, num)
		}
		fmt.Println("nums", nums)
		answer += solve(nums)
	}
	fmt.Println("answer", answer)
}

func secondPart(f *os.File) {
	scanner := bufio.NewScanner(f)
	answer := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		nums := make([]int, 0)
		for _, t := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(t)
			nums = append(nums, num)
		}
		fmt.Println("nums", nums)
		answer += solve2(nums)
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
