package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	data := readInput()

	parse(data)
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

type point struct {
	X  int
	Y  int
	Vx int
	Vy int
}

func min(n, v int) int {
	if n < v {
		return n
	}
	return v
}
func max(n, v int) int {
	if n > v {
		return n
	}
	return v
}

func parse(data []string) {

	points := []point{}
	for _, s := range data {
		p := point{}
		fmt.Sscanf(s, "position=<%d,%d> velocity=<%d, %d>", &p.X, &p.Y, &p.Vx, &p.Vy)
		points = append(points, p)
	}

	areamin := math.MaxInt64
	second := 0
	height := 0
	width := 0
	xOffset := 0
	yOffset := 0
	for i := 0; i < 15000; i++ {
		xMin := math.MaxInt64
		yMin := math.MaxInt64
		xMax := math.MinInt64
		yMax := math.MinInt64
		for _, p := range points {
			xMin = min(xMin, p.X+i*p.Vx)
			xMax = max(xMax, p.X+i*p.Vx)
			yMin = min(yMin, p.Y+i*p.Vy)
			yMax = max(yMax, p.Y+i*p.Vy)
		}
		if (xMax-xMin)*(yMax-yMin) < areamin {
			areamin = (xMax - xMin) * (yMax - yMin)
			height = yMax - yMin
			width = xMax - xMin
			xOffset = xMin
			yOffset = yMin
			second = i
		}
	}
	fmt.Printf("message readable at second %d\n", second)

	out := make([][]rune, width+1)
	for i := range out {
		out[i] = make([]rune, height+1)
	}
	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			out[x][y] = '.'
		}
	}
	for _, p := range points {
		x := (p.X + second*p.Vx) - xOffset
		y := (p.Y + second*p.Vy) - yOffset
		out[x][y] = 'X'
	}

	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			fmt.Printf("%c", out[x][y])
		}
		fmt.Printf("\n")
	}
}
