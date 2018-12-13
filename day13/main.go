package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// Cart current position, direction, crossroad turn counter
type Cart struct {
	pos       complex64
	direction complex64
	turns     int
}

// helpers for sorting carts based on position, top ones first
type cartSort []*complex64

func (cart cartSort) Len() int {
	return len(cart)
}

func (cart cartSort) Swap(i, j int) {
	cart[i], cart[j] = cart[j], cart[i]
}

func (cart cartSort) Less(i, j int) bool {
	return imag(*cart[i]) < imag(*cart[j]) || (imag(*cart[i]) == imag(*cart[j]) && real(*cart[i]) < real(*cart[j]))
}

func main() {
	var data []byte
	var maxWidth int

	data, maxWidth = readInput("input.txt")

	doStuff(data, maxWidth)
}

func readInput(filename string) (data []byte, w int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Bytes()
		if w == 0 {
			w = len(s)
		}
		data = append(data, s...)
	}

	return data, w
}

func doStuff(data []byte, maxWidth int) (float32, float32) {
	var cartPos []*complex64
	var carts = make(map[complex64]*Cart)
	var firstCrash complex64

	for i := 0; i < len(data); i++ {
		pos := complex(float32(i%maxWidth), float32(i/maxWidth))
		switch data[int(real(pos))+int(imag(pos))*maxWidth] {
		case '>':
			carts[pos] = &Cart{pos, 1, 0}
		case '<':
			carts[pos] = &Cart{pos, -1, 0}
		case 'v':
			carts[pos] = &Cart{pos, 1i, 0}
		case '^':
			carts[pos] = &Cart{pos, -1i, 0}
		default:
			continue
		}
		cartPos = append(cartPos, &carts[pos].pos)
	}

	for len(carts) > 1 {
		sort.Sort(cartSort(cartPos))

		for _, cart := range cartPos {
			if *cart != 0 {
				crashPos := carts[*cart].move(data, carts, maxWidth)
				if crashPos != complex(0, 0) && firstCrash == complex(0, 0) {
					firstCrash = crashPos
				}
			}
		}
	}

	fmt.Printf("First crash at %d,%d\n", int(real(firstCrash)), int(imag(firstCrash)))

	// hopefully a very short loop...
	for pos := range carts {
		fmt.Printf("Last cart standing %d,%d\n", int(real(pos)), int(imag(pos)))
	}

	return real(firstCrash), imag(firstCrash)
}

func (c *Cart) move(data []byte, carts map[complex64]*Cart, maxWidth int) (ret complex64) {
	delete(carts, c.pos)
	c.pos += c.direction
	if crash, ok := carts[c.pos]; ok {
		ret = c.pos
		delete(carts, c.pos)
		crash.pos = complex(0, 0)
		c.pos = complex(0, 0)
		return ret
	}
	carts[c.pos] = c

	if data[int(real(c.pos))+int(imag(c.pos))*maxWidth] == '/' {
		c.direction = complex(-imag(c.direction), -real(c.direction))
	}
	if data[int(real(c.pos))+int(imag(c.pos))*maxWidth] == '\\' {
		c.direction = complex(imag(c.direction), real(c.direction))
	}
	if data[int(real(c.pos))+int(imag(c.pos))*maxWidth] == '+' {
		c.turns++
		switch c.turns % 3 {
		case 0:
			c.direction *= 1i
		case 1:
			c.direction *= -1i
		case 2:
			c.direction *= 1
		}
	}

	return complex(0, 0)
}
