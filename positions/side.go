package positions

type Side struct {
	Name string
}

func (s *Side) MoveRight() {
	switch s.Name {
	case "N":
		s.Name = "I"
	case "I":
		s.Name = "S"
	case "S":
		s.Name = "W"
	case "W":
		s.Name = "N"
	}
}
func (s *Side) MoveLeft() {
	switch s.Name {
	case "N":
		s.Name = "W"
	case "I":
		s.Name = "N"
	case "S":
		s.Name = "I"
	case "W":
		s.Name = "S"
	}
}
