package main

import (
	"bufio"
	"fmt"
	"os"
)

type OpCode int

const (
	addr OpCode = iota
	addi
	mulr
	muli
	banr
	bani
	borr
	bori
	setr
	seti
	gtir
	gtri
	gtrr
	eqir
	eqri
	eqrr
)

func main() {
	var before [][4]int
	var after [][4]int
	var code [][4]int

	f, _ := os.Open("input1.txt")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		var tmp [4]int
		if len(s) > 6 {
			if s[0:6] == "Before" {
				fmt.Sscanf(s, "Before: [%d, %d, %d, %d]", &tmp[0], &tmp[1], &tmp[2], &tmp[3])
				before = append(before, tmp)
			} else if s[0:5] == "After" {
				fmt.Sscanf(s, "After: [%d, %d, %d, %d]", &tmp[0], &tmp[1], &tmp[2], &tmp[3])
				after = append(after, tmp)
			} else {
				fmt.Sscanf(s, "%d %d %d %d", &tmp[0], &tmp[1], &tmp[2], &tmp[3])
				code = append(code, tmp)
			}
		}
	}

	tot := 0
	ocmap := map[int]map[int]bool{}
	for i := 0; i < 16; i++ {
		ocmap[i] = map[int]bool{}
		for j := 0; j < 16; j++ {
			ocmap[i][j] = true
		}
	}
	for i := 0; i < len(before); i++ {
		matches := []int{}
		for opcode := OpCode(0); opcode < 16; opcode++ {
			ret := process(opcode, code[i][1], code[i][2], before[i])
			if ret == after[i][code[i][3]] {
				matches = append(matches, int(opcode))
			}
		}
		if len(matches) >= 3 {
			tot++
		}
	next:
		for ismatch := range ocmap[code[i][0]] {
			for _, v := range matches {
				if v == ismatch {
					continue next
				}
			}
			delete(ocmap[code[i][0]], ismatch)
		}
	}

	fmt.Printf("p1: %d\n", tot)

	oc := map[int]int{}
	for len(oc) < 16 {
		for i, matches := range ocmap {
			if len(matches) == 1 {
				for j := range matches {
					oc[i] = j
				}
				for j := range ocmap {
					delete(ocmap[j], oc[i])
				}
			}
		}
	}

	f, _ = os.Open("input2.txt")

	scanner = bufio.NewScanner(f)
	regs := [4]int{0, 0, 0, 0}
	for scanner.Scan() {
		s := scanner.Text()
		var tmp [4]int
		if len(s) > 6 {
			fmt.Sscanf(s, "%d %d %d %d", &tmp[0], &tmp[1], &tmp[2], &tmp[3])
			ret := process(OpCode(oc[tmp[0]]), tmp[1], tmp[2], regs)
			regs[tmp[3]] = ret
		}
	}
	fmt.Printf("p2: %d\n", regs[0])
}

func process(opcode OpCode, param1 int, param2 int, registers [4]int) int {
	switch opcode {
	case addr:
		return registers[param1] + registers[param2]
	case addi:
		return registers[param1] + param2
	case mulr:
		return registers[param1] * registers[param2]
	case muli:
		return registers[param1] * param2
	case banr:
		return registers[param1] & registers[param2]
	case bani:
		return registers[param1] & param2
	case borr:
		return registers[param1] | registers[param2]
	case bori:
		return registers[param1] | param2
	case setr:
		return registers[param1]
	case seti:
		return param1
	case gtir:
		if param1 > registers[param2] {
			return 1
		}
		return 0
	case gtri:
		if registers[param1] > param2 {
			return 1
		}
		return 0
	case gtrr:
		if registers[param1] > registers[param2] {
			return 1
		}
		return 0
	case eqir:
		if param1 == registers[param2] {
			return 1
		}
		return 0
	case eqri:
		if registers[param1] == param2 {
			return 1
		}
		return 0
	case eqrr:
		if registers[param1] == registers[param2] {
			return 1
		}
		return 0
	}
	return 0
}
