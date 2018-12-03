package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Claim struct {
	ID     string
	X      int
	Y      int
	Width  int
	Height int
}

func main() {
	data := readInput()

	claims := parseClaims(data)

	p1, p2 := getParts(claims)
	fmt.Printf("part 1: %d\npart 2: %s\n", p1, p2)
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

func parseClaims(data []string) (claims []Claim) {
	for i := range data {
		c := Claim{}
		fmt.Sscanf(data[i], "%s @ %d,%d: %dx%d", &c.ID, &c.X, &c.Y, &c.Width, &c.Height)
		claims = append(claims, c)
	}
	return claims
}

func getParts(claims []Claim) (total int, id string) {
	// assumptions: size < 1000, no more than one overlapping claim
	m := [1000][1000]int{}

	for _, c := range claims {
		for i := c.X; i < (c.X + c.Width); i++ {
			for j := c.Y; j < (c.Y + c.Height); j++ {
				m[i][j]++
			}
		}
	}
	for _, i := range m {
		for _, j := range i {
			if j > 1 {
				total++
			}
		}
	}

	for _, c := range claims {
		notOverlapping := true
		for i := c.X; i < (c.X + c.Width); i++ {
			for j := c.Y; j < (c.Y + c.Height); j++ {
				if m[i][j] > 1 {
					notOverlapping = false
				}
			}
		}
		if notOverlapping {
			id = c.ID
		}
	}
	return total, id
}
