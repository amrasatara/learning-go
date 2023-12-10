package aoc2023

import (
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
)

func Day10Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day10.txt")
	ParsePipes(input)
	return markPipes()
}

func Day10Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day10.txt")
	ParsePipes(input)
	markPipes()
	//clean
	for _, pipe := range pipes {
		if !pipe.Start && pipe.Marked == 0 {
			pipe.Value = '.'
		}
	}

	s := FindS()
	s1 := s.Next[0]

	AddToGroup(s, s1, input)
	for _, pd := range groupleft {
		if !pd.Start && pd.Marked == 0 {
			pd.Outside = "O"
			pd.MarkEmptyAround()
		}
	}

	for _, pd := range groupright {
		if !pd.Start && pd.Marked == 0 {
			pd.Outside = "I"
			pd.MarkEmptyAround()
		}
	}

	count := 0
	for _, pd := range pipes {
		if pd.Outside == "I" {
			count++
		}
	}

	return count
}

func AddToGroup(s *PipeDay10, s1 *PipeDay10, input []string) {

	if s.Y < s1.Y {
		// groupright = append(groupright, GetPipe(input, s.X-1, s.Y-1))
		groupright = append(groupright, GetPipe(input, s.X-1, s.Y))
		groupright = append(groupright, GetPipe(input, s.X-1, s.Y+1))

		// groupleft = append(groupleft, GetPipe(input, s.X+1, s.Y-1))
		groupleft = append(groupleft, GetPipe(input, s.X+1, s.Y))
		groupleft = append(groupleft, GetPipe(input, s.X+1, s.Y+1))
	} else if s.Y > s1.Y {
		// groupleft = append(groupleft, GetPipe(input, s.X-1, s.Y-1))
		groupleft = append(groupleft, GetPipe(input, s.X-1, s.Y))
		// groupleft = append(groupleft, GetPipe(input, s.X-1, s.Y+1))

		// groupright = append(groupright, GetPipe(input, s.X+1, s.Y-1))
		groupright = append(groupright, GetPipe(input, s.X+1, s.Y))
		// groupright = append(groupright, GetPipe(input, s.X+1, s.Y+1))

	} else if s.X < s1.X {

		// groupleft = append(groupleft, GetPipe(input, s.X-1, s.Y-1))
		groupleft = append(groupleft, GetPipe(input, s.X, s.Y-1))
		groupleft = append(groupleft, GetPipe(input, s.X+1, s.Y-1))

		// groupright = append(groupright, GetPipe(input, s.X-1, s.Y+1))
		groupright = append(groupright, GetPipe(input, s.X, s.Y+1))
		groupright = append(groupright, GetPipe(input, s.X+1, s.Y+1))
	} else if s.X > s1.X {

		// groupleft = append(groupleft, GetPipe(input, s.X-1, s.Y+1))
		groupleft = append(groupleft, GetPipe(input, s.X, s.Y+1))
		// groupleft = append(groupleft, GetPipe(input, s.X+1, s.Y+1))

		// groupright = append(groupright, GetPipe(input, s.X-1, s.Y-1))
		groupright = append(groupright, GetPipe(input, s.X, s.Y-1))
		// groupright = append(groupright, GetPipe(input, s.X+1, s.Y-1))
	}

	next := s1.FindNext(s)
	if next.Start {
		return
	}
	AddToGroup(s1, next, input)
}

var groupleft []*PipeDay10
var groupright []*PipeDay10

func markPipes() int {
	s := FindS()
	s.Marked = 0

	next := make([]*PipeDay10, len(s.Next))
	copy(next, s.Next)
	steps := 1
	for ; ; steps++ {
		nextNext := make([]*PipeDay10, 0)
		for _, pd := range next {
			pd.Marked = steps
			for _, pd2 := range pd.Next {
				if !pd2.Start && pd2.Marked == 0 {
					nextNext = append(nextNext, pd2)
				}
			}
		}
		if len(nextNext) == 0 {
			break
		}
		next = nextNext

	}
	return steps
}

var pipes []*PipeDay10

func ParsePipes(input []string) {
	for i, v := range input {
		for j := range v {
			p := GetPipe(input, j, i)
			if p == nil {
				continue
			}
			p.Left = GetPipe(input, j, i-1)
			p.Right = GetPipe(input, j, i+1)
			p.Down = GetPipe(input, j+1, i)
			p.Up = GetPipe(input, j-1, i)
		}
	}
	s := FindS()

	RecalculateForS(s)

	for _, p := range pipes {
		FindLinks(p)
	}
}

func RecalculateForS(p *PipeDay10) {
	var r, l, d, u bool
	if p.Right != nil && HasAnyValue(p.Right, []rune{'-', '7', 'J'}) {
		p.Next = append(p.Next, p.Right)
		r = true
	}
	if p.Down != nil && HasAnyValue(p.Down, []rune{'|', 'L', 'J'}) {
		p.Next = append(p.Next, p.Down)
		d = true
	}
	if p.Left != nil && HasAnyValue(p.Left, []rune{'-', 'L', 'F'}) {
		p.Next = append(p.Next, p.Left)
		l = true
	}
	if p.Up != nil && HasAnyValue(p.Up, []rune{'|', '7', 'F'}) {
		p.Next = append(p.Next, p.Up)
		u = true
	}

	if u && d {
		p.Value = '|'
	} else if u && r {
		p.Value = 'L'
	} else if u && l {
		p.Value = 'J'
	} else if d && r {
		p.Value = 'F'
	} else if d && l {
		p.Value = '7'
	} else if r && l {
		p.Value = '-'
	}
}
func FindS() *PipeDay10 {
	for _, r := range pipes {
		if r.Start {
			return r
		}
	}

	return nil
}

func FindPipe(x, y int) *PipeDay10 {
	for _, r := range pipes {
		if r.Y == y && r.X == x {
			return r
		}
	}

	return nil
}
func FindLinks(p *PipeDay10) {
	p.Next = make([]*PipeDay10, 0, 4)
	switch p.Value {
	case 'F':
		if p.Right != nil {
			p.Next = append(p.Next, p.Right)
		}
		if p.Down != nil {
			p.Next = append(p.Next, p.Down)
		}
	case '|':
		if p.Up != nil {
			p.Next = append(p.Next, p.Up)
		}
		if p.Down != nil {
			p.Next = append(p.Next, p.Down)
		}
	case '-':
		if p.Left != nil {
			p.Next = append(p.Next, p.Left)
		}
		if p.Right != nil {
			p.Next = append(p.Next, p.Right)
		}
	case '7':
		if p.Left != nil {
			p.Next = append(p.Next, p.Left)
		}
		if p.Down != nil {
			p.Next = append(p.Next, p.Down)
		}
	case 'L':
		if p.Up != nil {
			p.Next = append(p.Next, p.Up)
		}
		if p.Right != nil {
			p.Next = append(p.Next, p.Right)
		}
	case 'J':
		if p.Up != nil {
			p.Next = append(p.Next, p.Up)
		}
		if p.Left != nil {
			p.Next = append(p.Next, p.Left)
		}
	}
}

func HasAnyValue(pipe *PipeDay10, list []rune) bool {
	for _, v := range list {
		if pipe.Value == v {
			return true
		}
	}

	return false
}

func GetPipe(input []string, x int, y int) *PipeDay10 {
	maxX := len(input)
	maxY := len(input[0])
	if x < 0 || y < 0 || x >= maxX || y >= maxY {
		return nil
	}

	value := rune(input[x][y])

	for _, r := range pipes {
		if r.X == x && r.Y == y {
			return r
		}
	}
	r2 := PipeDay10{
		Value: value,
		X:     x,
		Y:     y,
		Start: value == 'S',
	}
	pipes = append(pipes, &r2)

	return &r2

}

type PipeDay10 struct {
	Value   rune
	X       int
	Y       int
	Start   bool
	Left    *PipeDay10
	Right   *PipeDay10
	Up      *PipeDay10
	Down    *PipeDay10
	Next    []*PipeDay10
	Marked  int
	Outside string
}

func (p *PipeDay10) FindNext(from *PipeDay10) *PipeDay10 {

	for _, pd := range p.Next {
		if pd != from {
			return pd
		}
	}

	return nil
}

func (p *PipeDay10) IsEmpty() bool {
	x := p != nil && !p.Start && p.Marked == 0
	return x
}

func (p *PipeDay10) MarkEmptyAround() {
	if p.Left != nil && p.Left.IsEmpty() && len(p.Left.Outside) == 0 {
		p.Left.Outside = p.Outside
		p.Left.MarkEmptyAround()
	}
	if p.Right != nil && p.Right.IsEmpty() && len(p.Right.Outside) == 0 {
		p.Right.Outside = p.Outside
		p.Right.MarkEmptyAround()
	}
	if p.Up != nil && p.Up.IsEmpty() && len(p.Up.Outside) == 0 {
		p.Up.Outside = p.Outside
		p.Up.MarkEmptyAround()
	}
	if p.Down != nil && p.Down.IsEmpty() && len(p.Down.Outside) == 0 {
		p.Down.Outside = p.Outside
		p.Down.MarkEmptyAround()
	}
}
func (p PipeDay10) ToString() string {
	sb := strings.Builder{}

	sb.WriteRune(p.Value)
	if p.Left != nil {
		sb.WriteString(" Left: ")
		sb.WriteRune(p.Left.Value)
	}

	if p.Right != nil {
		sb.WriteString(" Right: ")
		sb.WriteRune(p.Right.Value)
	}

	if p.Down != nil {
		sb.WriteString(" Down: ")
		sb.WriteRune(p.Down.Value)
	}

	if p.Up != nil {
		sb.WriteString(" Up: ")
		sb.WriteRune(p.Up.Value)
	}
	return sb.String()
}

func printInFile(input []string) {
	sb := strings.Builder{}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			p := FindPipe(j, i)
			if len(p.Outside) > 0 {
				sb.WriteString(p.Outside)
			} else if !p.Start && p.Marked == 0 {
				sb.WriteString(" ")
			} else if p.Start {
				sb.WriteString(" ")
			} else if p.Marked > 0 {
				sb.WriteString("#")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	fileparsers.Write(sb, "temp.txt")
}
