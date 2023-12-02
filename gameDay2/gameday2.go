package gameday2

func (game *GameDay2) CountRed() int {
	result := 0

	for _, r := range game.Rounds {
		result += r.Red
	}

	return result
}
func (game *GameDay2) CountGreen() int {
	result := 0

	for _, r := range game.Rounds {
		result += r.Green
	}

	return result
}
func (game *GameDay2) CountBlue() int {
	result := 0

	for _, r := range game.Rounds {
		result += r.Blue
	}

	return result
}

func (game *GameDay2) MaxRed() int {
	result := 0

	for _, r := range game.Rounds {
		if r.Red > result {
			result = r.Red
		}
	}

	return result
}
func (game *GameDay2) MaxGreen() int {
	result := 0

	for _, r := range game.Rounds {
		if r.Green > result {
			result = r.Green
		}
	}

	return result
}
func (game *GameDay2) MaxBlue() int {
	result := 0

	for _, r := range game.Rounds {
		if r.Blue > result {
			result = r.Blue
		}
	}

	return result
}

type GameDay2 struct {
	Id     int
	Rounds []GameRoundDay2
}

type GameRoundDay2 struct {
	Green int
	Red   int
	Blue  int
}
