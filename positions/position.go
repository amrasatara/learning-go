package positions

import (
	"fmt"
	"math"
)

type Position struct {
	X    int
	Y    int
	Side Side
}

func (p *Position) CalculateFrom0() int {
	return int(math.Abs(float64(p.X)-0) + math.Abs(float64(p.Y)-0))
}

func (p *Position) Move(RL string, steps int, visited map[string]bool) bool {
	switch RL {
	case "R":
		p.Side.MoveRight()
	case "L":
		p.Side.MoveLeft()
	}

	for i := 0; i < steps; i++ {
		key := p.UniqueStringName()
		if visited != nil {
			_, ok := visited[key]
			if ok {
				return true
			}
			visited[key] = true
		}
		switch p.Side.Name {
		case "N":
			p.Y++
		case "S":
			p.Y--
		case "I":
			p.X++
		case "W":
			p.X--
		}
	}
	return false
}

func (p *Position) UniqueStringName() string {
	return fmt.Sprintf("%d|%d", p.X, p.Y)
}
