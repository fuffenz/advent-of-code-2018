package main

import "fmt"

func main() {

	//parse(18)
	//parse(42)
	parse(3031)
}

func parse(serial int) {

	fmt.Printf("Input serial: %d\n", serial)
	grid := [301][301]int{}

	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			id := x + 10
			power := ((((id*y + serial) * id) / 100) % 10) - 5
			grid[x][y] = power
		}
	}

	totalMax := 0
	totalX := 0
	totalY := 0
	best := 0
	for size := 3; size < 300; size++ {
		for x := 1; x <= (300 - size); x++ {
			for y := 1; y <= (300 - size); y++ {
				total := 0
				for x1 := 0; x1 < size; x1++ {
					for y1 := 0; y1 < size; y1++ {
						total += grid[x+x1][y+y1]
					}
				}
				if total > totalMax {
					totalMax = total
					totalX = x
					totalY = y
				}
			}
		}
		//fmt.Printf("Size %d, largest %d (%d, %d)\n", size, totalMax, totalX, totalY)
		if totalMax > best {
			fmt.Printf("Best so far: %d  (%d, %d) size %d \n", totalMax, totalX, totalY, size)
			best = totalMax
		}
	}
}
