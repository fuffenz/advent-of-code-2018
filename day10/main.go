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

	area := math.MaxInt64
	prevArea := math.MaxInt64
	second := 0
	height := 0
	width := 0
	xOffset := 0
	yOffset := 0
	for {
		xMin := math.MaxInt64
		yMin := math.MaxInt64
		xMax := math.MinInt64
		yMax := math.MinInt64
		for _, p := range points {
			xMin = min(xMin, p.X+second*p.Vx)
			xMax = max(xMax, p.X+second*p.Vx)
			yMin = min(yMin, p.Y+second*p.Vy)
			yMax = max(yMax, p.Y+second*p.Vy)
		}
		area = (xMax - xMin) * (yMax - yMin)
		if area < prevArea {
			height = yMax - yMin
			width = xMax - xMin
			xOffset = xMin
			yOffset = yMin
		}
		if area > prevArea {
			break
		}
		prevArea = area
		second++
	}
	second--
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
