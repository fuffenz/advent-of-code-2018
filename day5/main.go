package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")

	p1, p2 := getResults(string(data))

	fmt.Printf("Part 1 %d, part 2 %d", p1, p2)
}

func reactPair(s string, unitpairs []string) string {
	for _, u := range unitpairs {
		s = strings.Replace(s, u, "", -1)
	}
	return s
}

func reactAll(s string, unitpairs []string) string {
	tmp := s
	for true {
		res := reactPair(tmp, unitpairs)
		if res == tmp {
			break
		}
		tmp = res
	}
	return tmp
}

func getResults(s string) (int, int) {
	var unitpairs []string
	for c := 'A'; c <= 'Z'; c++ {
		unitpairs = append(unitpairs, fmt.Sprintf("%c%c", c, c+32))
		unitpairs = append(unitpairs, fmt.Sprintf("%c%c", c+32, c))
	}

	// part 1
	tmp := reactAll(s, unitpairs)
	p1 := len(tmp)

	// part 2
	ss := s
	minLen := len(tmp)
	for c := 'A'; c <= 'Z'; c++ {
		tmp := strings.Replace(ss, string(c), "", -1)
		tmp = strings.Replace(tmp, string(c+32), "", -1)
		tmp = reactAll(tmp, unitpairs)
		if len(tmp) < minLen {
			minLen = len(tmp)
		}
	}

	return p1, minLen
}
