package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data := readInput()

	getResult(data)
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
func nextGeneration(state string, rules map[string]string, iter int) (string, int) {
	slask := "..." + state + "....."
	out := []rune(state + ".....")
	for i := 0; i < len(state); i++ {
		tmp := slask[i : i+5]
		didMatch := false
		for pattern, result := range rules {
			if tmp == pattern {
				out[i+2] = rune(result[0])
				didMatch = true
			}
		}
		if didMatch == false {
			out[i+2] = '.'
		}
	}

	s := string(out)
	tot := 0
	for i := 0; i < len(s); i++ {
		pos := i - (iter * 3) - 6
		if s[i] == '#' {
			tot += pos
		}
	}
	return string(s), tot
}

func getResult(data []string) {

	var state string
	fmt.Sscanf(data[0], "initial state: %s", &state)

	state = "..." + state + "....."
	var rules = make(map[string]string)
	for i := 2; i < len(data); i++ {
		var s string
		var s2 string
		fmt.Sscanf(data[i], "%s => %s", &s, &s2)
		rules[s] = s2
	}

	s := state
	var c int
	for i := 0; i < 20; i++ {
		s, c = nextGeneration(s, rules, i)
	}
	fmt.Printf("p1: %d\n", c)

	s = state
	var diff, last int
	for i := 0; i < 200; i++ {
		s, c = nextGeneration(s, rules, i)
		diff = c - last
		last = c
	}

	p2 := (uint64(50000000000)-200)*uint64(diff) + uint64(last)
	fmt.Printf("p2: %d\n", p2)
}
