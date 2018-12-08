package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("input.txt")
	s := strings.Split(string(b), " ")
	data := make([]int, len(s))

	for i := range data {
		data[i], _ = strconv.Atoi(s[i])
	}

	p1, p2, _ := getResults(data)

	fmt.Printf("Part 1 %d, part 2 %d", p1, p2)
}

func getResults(data []int) (int, int, []int) {
	childnodes := data[0]
	metacount := data[1]
	rest := data[2:]
	values := make(map[int]int)
	total := 0

	for i := 0; i < childnodes; i++ {
		var t, v int
		t, v, rest = getResults(rest)
		total += t
		values[i] = v
	}

	metasum := 0
	for i := 0; i < metacount; i++ {
		metasum += rest[i]
	}
	total += metasum

	if childnodes == 0 {
		return total, metasum, rest[metacount:]
	}

	totvalue := 0
	for i := 0; i < metacount; i++ {
		v := rest[i]
		if _, ok := values[v-1]; ok {
			totvalue += values[v-1]
		}
	}
	return total, totvalue, rest[metacount:]

}
