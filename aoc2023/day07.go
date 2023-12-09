package aoc2023

import (
	"fmt"
	"sort"

	"github.com/amra.satara/learning-go/fileparsers"
)

func Day07Part1() int {
	input := fileparsers.ReadLines("inputs2023\\day7.txt")

	hands := make([]HandDay7, len(input))
	for i, v := range input {
		h := ParseHand(v, false)
		hands[i] = h
	}

	result := SortAndCalculate(hands, false)
	return result
}

func Day07Part2() int {
	input := fileparsers.ReadLines("inputs2023\\day7.txt")

	hands := make([]HandDay7, len(input))
	for i, v := range input {
		h := ParseHand(v, true)
		hands[i] = h
	}

	result := SortAndCalculate(hands, true)
	return result
}

func SortAndCalculate(hands []HandDay7, joker bool) int {
	sort.Slice(hands, func(i, j int) bool {
		h1 := hands[i]
		h2 := hands[j]
		if h1.Score < h2.Score {
			return true
		}
		if h1.Score > h2.Score {
			return false
		}
		for i := 0; i < len(h1.CardsRaw); i++ {
			r1 := []rune(h1.CardsRaw)[i]
			r2 := []rune(h2.CardsRaw)[i]

			if r1 == r2 {
				continue
			}

			if joker {
				if r1 == 'J' {
					return true
				}
				if r2 == 'J' {
					return false
				}
			}

			if IndexOf(cardTypes, r1) > IndexOf(cardTypes, r2) {
				return true
			} else {
				return false
			}
		}
		return true
	})

	result := 0

	for i, hd := range hands {
		result += (i + 1) * hd.Bid
	}

	return result
}

func IndexOf(list []rune, v rune) int {
	for i, v2 := range list {
		if v == v2 {
			return i
		}
	}
	return -1
}

var cardTypes = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

func ParseHand(input string, JAsJoker bool) HandDay7 {
	result := HandDay7{}

	fmt.Sscanf(input, "%s %d", &result.CardsRaw, &result.Bid)

	result.CardDict = make(map[rune]int)
	result.NumberOfJokers = CountR(result.CardsRaw, 'J')
	for _, r := range cardTypes {
		c := CountR(result.CardsRaw, r)
		if c == 0 {
			continue
		}
		result.CardDict[r] = c
		result.Score = CalculateScore(&result, c)
	}

	if result.Score == 7 || !JAsJoker || result.NumberOfJokers == 0 {
		return result
	}

	numb := make([]int, 0)
	for k, v := range result.CardDict {
		if k != 'J' {
			numb = append(numb, v)
		}
	}

	sort.Slice(numb, func(i, j int) bool {
		return numb[i] < numb[j]
	})

	switch len(numb) {
	case 1:
		result.Score = 7
	case 4:
		result.Score = 2
	case 3:
		result.Score = 4
	case 2:
		if numb[0] == 1 {
			result.Score = 6
		} else {
			result.Score = 5
		}

	}

	return result

}

func CalculateScoreWithJoker(result *HandDay7, c int) int {
	if c == 5 {
		return 7
	} else if c == 4 {
		return 6
	} else if c == 3 {
		if result.Score == 2 {
			return 5 // fullhouse
		} else {
			return 4
		}
	} else if c == 2 {
		if result.Score == 4 {
			return 5
		} else if result.Score == 2 {
			return 3
		} else {
			return 2
		}
	} else if c == 1 && result.Score == 0 {
		return 1
	}
	return result.Score
}

func CalculateScore(result *HandDay7, c int) int {
	if c == 5 {
		return 7
	} else if c == 4 {
		return 6
	} else if c == 3 {
		if result.Score == 2 {
			return 5 // fullhouse
		} else {
			return 4
		}
	} else if c == 2 {
		if result.Score == 4 {
			return 5
		} else if result.Score == 2 {
			return 3
		} else {
			return 2
		}
	} else if c == 1 && result.Score == 0 {
		return 1
	}
	return result.Score
}

// 7 Five of a kind, where all five cards have the same label: AAAAA
// 6 Four of a kind, where four cards have the same label and one card has a different label: AA8AA
// 5 Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
// 4 Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
// 3 Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
// 2 One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
// 1 High card, where all cards' labels are distinct: 23456

func CountR(list string, tr rune) int {
	result := 0
	for _, v := range list {
		if v == tr {
			result++
		}
	}

	return result
}

type HandDay7 struct {
	CardsRaw       string
	Bid            int
	Score          int
	CardDict       map[rune]int
	NumberOfJokers int
}
