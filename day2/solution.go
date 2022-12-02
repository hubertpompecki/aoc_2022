package day2

import (
	"strings"

	"github.com/hpompecki/aoc_2022/input"
)

func TotalPoints(fileName string) int {
	var score int
	for round := range input.Read(fileName) {
		score += calculateScore(round)
	}
	return score
}

func TotalPoints2(fileName string) int {
	var score int
	for round := range input.Read(fileName) {
		score += calculateScore2(round)
	}
	return score
}

func calculateScore(r string) int {
	var score int
	round := strings.Split(r, " ")

	// Points for our choice
	switch round[1] {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	}

	// Points for result
	switch {
	case round[0] == "A" && round[1] == "X":
		score += 3
	case round[0] == "B" && round[1] == "Y":
		score += 3
	case round[0] == "C" && round[1] == "Z":
		score += 3
	case round[0] == "A" && round[1] == "Y":
		score += 6
	case round[0] == "B" && round[1] == "Z":
		score += 6
	case round[0] == "C" && round[1] == "X":
		score += 6
	default:
	}

	return score
}

func calculateScore2(r string) int {
	var score int
	round := strings.Split(r, " ")

	// Points for result
	switch round[1] {
	case "X":
		score += 0
	case "Y":
		score += 3
	case "Z":
		score += 6
	}

	// Points for suggested choice
	switch {
	case round[0] == "A" && round[1] == "X":
		score += 3
	case round[0] == "B" && round[1] == "X":
		score += 1
	case round[0] == "C" && round[1] == "X":
		score += 2
	case round[0] == "A" && round[1] == "Y":
		score += 1
	case round[0] == "B" && round[1] == "Y":
		score += 2
	case round[0] == "C" && round[1] == "Y":
		score += 3
	case round[0] == "A" && round[1] == "Z":
		score += 2
	case round[0] == "B" && round[1] == "Z":
		score += 3
	case round[0] == "C" && round[1] == "Z":
		score += 1
	}

	return score
}
