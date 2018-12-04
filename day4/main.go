package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	data := readInput()

	p1, p2 := getResults(data)
	fmt.Printf("Part 1 %d, part 2 %d", p1, p2)
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

type guard struct {
	minutes [60]int
	total   int
}

func getResults(data []string) (int, int) {
	sort.Strings(data)

	guards := make(map[int]guard)
	guardID := 0
	var fellAsleep time.Time

	for _, s := range data {
		t, _ := time.Parse("2006-01-02 15:04", s[1:17])
		txt := s[19:]

		switch txt {
		case "falls asleep":
			fellAsleep = t
		case "wakes up":
			sleepDiff := t.Sub(fellAsleep)
			tot := int(sleepDiff.Minutes())
			g := guard{}
			if v, ok := guards[guardID]; ok {
				g = v
			}
			g.total += tot
			for i := fellAsleep.Minute(); i < t.Minute(); i++ {
				g.minutes[i]++
			}
			guards[guardID] = g
		default:
			fmt.Sscanf(txt, "Guard #%d", &guardID)
		}
	}

	maxMin := 0
	worstGuard := 0
	for g, v := range guards {
		if v.total > maxMin {
			worstGuard = g
			maxMin = v.total
		}
	}

	worstMinute := 0
	for m, v := range guards[worstGuard].minutes {
		if v > guards[worstGuard].minutes[worstMinute] {
			worstMinute = m
		}
	}
	part1 := worstGuard * worstMinute

	part2 := 0
	maxMinutes := 0
	for g, v := range guards {
		for i, minutes := range v.minutes {
			if minutes > maxMinutes {
				maxMinutes = minutes
				part2 = g * i
			}
		}
	}

	return part1, part2
}
