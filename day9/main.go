package main

import (
	"fmt"
)

type Marble struct {
	Next     *Marble
	Previous *Marble
	Value    int
}

func NewMarble(value int) *Marble {
	m := &Marble{
		Value: value,
	}
	m.Next = m
	m.Previous = m
	return m
}

func (m *Marble) AddMarble(newMarble *Marble) {
	currentNext := m.Next
	currentNext.Previous = newMarble
	newMarble.Previous = m
	newMarble.Next = currentNext
	m.Next = newMarble
}

func (m *Marble) RemoveNext() *Marble {
	currentNext := m.Next
	m.Next = currentNext.Next
	m.Next.Previous = m
	return currentNext
}

func main() {
	// 439 players; last marble is worth 71307 points

	p1 := WinningScore(439, 71307)

	p2 := WinningScore(439, 71307*100)

	fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

func WinningScore(numberOfPlayers int, numberOfMarbles int) int {
	players := make([]int, numberOfPlayers)
	marbleValue := 0
	highScore := 0
	currentMarble := NewMarble(0)
	for i := 0; i < numberOfMarbles; i++ {
		p := i % numberOfPlayers
		marbleValue++
		if marbleValue%23 > 0 {
			currentMarble = currentMarble.Next
			m := NewMarble(marbleValue)
			currentMarble.AddMarble(m)
			currentMarble = m
		} else {
			for j := 0; j < 8; j++ {
				currentMarble = currentMarble.Previous
			}
			removedMarble := currentMarble.RemoveNext()
			players[p] += (marbleValue + removedMarble.Value)
			if players[p] > highScore {
				highScore = players[p]
			}
			currentMarble = currentMarble.Next
		}
	}

	return highScore
}
