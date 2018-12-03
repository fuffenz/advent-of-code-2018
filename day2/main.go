package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data := readInput()

	fmt.Printf("part 1: %d\n", getCheckSum(data))
	fmt.Printf("part 2: %s\n", getSmallestDiff(data))
}

func readInput() (data []string) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}

func getCheckSum(data []string) int {

	twoCount := 0
	threeCount := 0

	for i := range data {
		isTwo := false
		isThree := false
		for _, char := range data[i] {
			count := strings.Count(data[i], string(char))
			if count == 2 {
				isTwo = true
			}
			if count == 3 {
				isThree = true
			}
		}
		if isTwo {
			twoCount++
		}
		if isThree {
			threeCount++
		}
	}
	return twoCount * threeCount
}

func mismatchCount(s1 string, s2 string) int {
	c := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			c++
		}
	}
	return c
}

func getSmallestDiff(data []string) string {

	for i := range data {
		for j := 0; j < i; j++ {
			if mismatchCount(data[i], data[j]) == 1 {
				res := ""
				for n := range data[i] {
					if data[i][n] == data[j][n] {
						res += string(data[i][n])
					}
				}
				return res
			}
		}
	}

	return ""
}
