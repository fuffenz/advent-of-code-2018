package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// xMin 414, xMax 576, yMin 8, yMax 1925
var grid [600][2000]byte

func main() {
	f, _ := os.Open("input.txt")

	var xMin, xMax int = math.MaxInt64, math.MinInt64
	var yMin, yMax int = math.MaxInt64, math.MinInt64

	for x := 0; x < 600; x++ {
		for y := 0; y < 2000; y++ {
			grid[x][y] = '.'
		}
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		var x1, x2, y1, y2 int
		if len(s) > 6 {
			if s[0] == 'x' { // x=426, y=960..968
				fmt.Sscanf(s, "x=%d, y=%d..%d", &x1, &y1, &y2)
				if x1 < xMin {
					xMin = x1
				}
				if x1 > xMax {
					xMax = x1
				}
				if y1 < yMin {
					yMin = y1
				}
				if y2 > yMax {
					yMax = y2
				}
				for y := y1; y <= y2; y++ {
					grid[x1][y] = '#'
				}
			}
			if s[0] == 'y' { // y=24, x=519..531
				fmt.Sscanf(s, "y=%d, x=%d..%d", &y1, &x1, &x2)
				if y1 < yMin {
					yMin = y1
				}
				if y1 > yMax {
					yMax = y1
				}
				if x1 < xMin {
					xMin = x1
				}
				if x2 > xMax {
					xMax = x2
				}
				for x := x1; x <= x2; x++ {
					grid[x][y1] = '#'
				}
			}
		}
	}

	//fmt.Printf("xMin %d, xMax %d, yMin %d, yMax %d\n", xMin, xMax, yMin, yMax)

	justAddWater(500, 0, yMax) // initial spring

	for y := yMin - 1; y <= yMax+1; y++ {
		for x := xMin - 1; x <= xMax+1; x++ {
			fmt.Printf("%c", grid[x][y])
		}
		fmt.Printf("\n")
	}

	rest := 0
	reach := 0
	for x := xMin - 1; x <= xMax+1; x++ {
		for y := yMin; y <= yMax; y++ {
			if grid[x][y] == '~' {
				rest++
			}
			if grid[x][y] == '|' {
				reach++
			}
		}
	}

	fmt.Printf("P1: %d\n", rest+reach)
	fmt.Printf("P2: %d\n", rest)
}

func justAddWater(x int, y int, yMax int) {
	if !canFlow(x, y) {
		return
	}

	if y > yMax {
		return
	}

	// down available?
	if canFlow(x, y+1) {
		if grid[x][y] == '.' {
			grid[x][y] = '|'
			justAddWater(x, y+1, yMax)
			if grid[x][y+1] == '~' {
				justAddWater(x, y, yMax)
			}
		}
	} else { // not down, but maybe left/right
		left := x
		for canFlow(left, y) && !canFlow(left, y+1) {
			grid[left][y] = '|'
			left--
		}
		right := x + 1
		for canFlow(right, y) && !canFlow(right, y+1) {
			grid[right][y] = '|'
			right++
		}
		if canFlow(left, y+1) || canFlow(right, y+1) {
			justAddWater(left, y, yMax)
			justAddWater(right, y, yMax)
		} else if grid[left][y] == '#' && grid[right][y] == '#' {
			for i := left + 1; i < right; i++ {
				grid[i][y] = '~'
			}
		}
	}
}

func canFlow(x int, y int) bool {
	return grid[x][y] == '.' || grid[x][y] == '|'
}
