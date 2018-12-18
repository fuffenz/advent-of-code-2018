package main

import (
	"bufio"
	"fmt"
	"os"
)

var grid [50][]byte

const gridsize = 50

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		s := scanner.Text()
		grid[i] = []byte(s)
		i++
	}

	backlog := []int{}
	patternStart := 0
	prevSum := 0
	tickP2 := 1000000000
	for i = 1; i < 1000 && patternStart == 0; i++ {
		trees, lumberyards := tick()
		sum := trees * lumberyards
		if i == 10 {
			fmt.Printf("P1: tick %d : %d trees, %d lumberyards = %d\n", i, trees, lumberyards, sum)
		}
		for j := 1; j < len(backlog); j++ {
			if backlog[j-1] == prevSum && backlog[j] == sum {
				fmt.Printf("pattern found %d\n", j)
				patternStart = j
				break
			}
		}
		backlog = append(backlog, sum)
		prevSum = sum
	}

	diff := i - patternStart - 2
	pos := (tickP2 - patternStart - 1) % diff

	fmt.Printf("p2: tick %d : score %d\n", tickP2, backlog[patternStart+pos])
}

func tick() (int, int) {
	var newgrid [50][]byte
	for i := 0; i < gridsize; i++ {
		newgrid[i] = make([]byte, gridsize)
		for j := 0; j < gridsize; j++ {
			trees, lumberyards := check(i, j)
			switch grid[i][j] {
			case '.':
				if trees > 2 {
					newgrid[i][j] = '|'
				} else {
					newgrid[i][j] = '.'
				}
			case '|':
				if lumberyards > 2 {
					newgrid[i][j] = '#'
				} else {
					newgrid[i][j] = '|'
				}
			case '#':
				if lumberyards < 1 || trees < 1 {
					newgrid[i][j] = '.'
				} else {
					newgrid[i][j] = '#'
				}
			}
		}
	}
	trees, lumberyards := 0, 0
	for i := 0; i < gridsize; i++ {
		for j := 0; j < gridsize; j++ {
			grid[i][j] = newgrid[i][j]
			switch grid[i][j] {
			case '|':
				trees++
			case '#':
				lumberyards++
			}

		}
	}
	return trees, lumberyards
}

func check(x, y int) (trees int, lumberyards int) {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if x+i >= 0 && x+i < gridsize && y+j >= 0 && y+j < gridsize && !(i == 0 && j == 0) {
				if grid[x+i][y+j] == '|' {
					trees++
				}
				if grid[x+i][y+j] == '#' {
					lumberyards++
				}
			}
		}
	}
	return trees, lumberyards
}
