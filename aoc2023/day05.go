package aoc2023

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/amra.satara/learning-go/fileparsers"
)

func Day05Part1() int64 {
	input := fileparsers.ReadLines("inputs2023\\day5.txt")

	seeds := TakeSeeds(input[0])

	mappers := ParseAsList(input)

	var minLoc int64 = 3952138264
	for _, seed := range seeds {

		for _, mapper := range mappers {
			seed, _ = FindDesc(mapper, seed)
		}
		minLoc = min(minLoc, seed)
	}

	return minLoc
}

func Day05Part2() int64 {
	input := fileparsers.ReadLines("inputs2023\\day5.txt")

	seeds := TakeSeeds(input[0])

	seedRanges := make([]SeedRangeDay5, 0, len(seeds)/2)

	for i := 0; i < len(seeds)-1; i += 2 {
		from := seeds[i]
		to := seeds[i] + seeds[i+1]
		seedRange := SeedRangeDay5{from, to}
		seedRanges = append(seedRanges, seedRange)
	}

	mappers := ParseAsList(input)

	for _, v := range mappers {
		Sort(v)
	}
	var j int64
	for j = 0; ; j++ {
		x := int64(j)

		for i := len(mappers) - 1; i >= 0; i-- {
			mapper := mappers[i]
			x = FindSource(mapper, x)
		}
		match := HasSeed(&seedRanges, x)
		//fmt.Println(x)
		if match {
			return j
		}
	}
}

func HasSeed(seedRanges *[]SeedRangeDay5, x int64) bool {
	for _, srd := range *seedRanges {
		if srd.From <= x && x < srd.To {
			return true
		}
	}
	return false
}

func Sort(list []MapLineDay5) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Destination < list[j].Destination
	})
}

func ParseAsList(input []string) [][]MapLineDay5 {
	result := make([][]MapLineDay5, 0)

	var current *[]MapLineDay5
	for _, v := range input[1:] {
		if len(v) == 0 {
			continue
		}
		if strings.Contains(v, " map:") {
			result = append(result, make([]MapLineDay5, 0))
			current = &result[len(result)-1]
			continue
		}
		line := MapLineDay5{}
		fmt.Sscanf(v, "%d%d%d", &line.Destination, &line.Source, &line.RangeLength)
		line.MaxSource = line.Source + line.RangeLength
		line.Dinst = line.Source - line.Destination
		line.MaxDestination = line.Destination + line.RangeLength
		*current = append(*current, line)

	}
	return result
}

func TakeSeeds(input string) []int64 {

	input = strings.ReplaceAll(input, "seeds: ", "")
	parts := strings.Split(input, " ")
	seeds := make([]int64, len(parts))
	for i, v := range parts {
		seeds[i], _ = strconv.ParseInt(v, 10, 64)
	}

	return seeds
}

func MakeMapper(list []string) []MapLineDay5 {

	result := make([]MapLineDay5, 0)

	return result
}

func FindDesc(mapper []MapLineDay5, input int64) (int64, *MapLineDay5) {
	for _, mline := range mapper {
		v, ok := mline.Find(input)
		if ok {
			return v, &mline
		}
	}

	return input, nil
}

// mapper is sorted
func FindSource(mapper []MapLineDay5, input int64) int64 {
	for _, mline := range mapper {
		if input >= mline.MaxDestination {
			continue
		}
		if input < mline.Destination || input >= mline.MaxDestination {
			continue
		}
		return input + mline.Dinst
	}

	return input
}

type MapLineDay5 struct {
	Destination    int64
	Source         int64
	RangeLength    int64
	MaxSource      int64
	MaxDestination int64
	Dinst          int64
}

func (mapline *MapLineDay5) Find(input int64) (int64, bool) {

	if input < mapline.Source || input > (mapline.MaxSource) {
		return input, false
	}

	return input - mapline.Dinst, true
}

type SeedRangeDay5 struct {
	From int64
	To   int64
}
