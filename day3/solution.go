package day3

import (
	"github.com/hpompecki/aoc_2022/input"
	"golang.org/x/exp/slices"
)

func SumOfPriorities(fileName string) int {
	var s int
	for contents := range input.Read(fileName) {
		d := findDuplicateItem(contents)
		s += priority(d)
	}
	return s
}

func SumOfBadgePriorities(fileName string) int {
	var s int
	var currentGroup = []string{}
	for elf := range input.Read(fileName) {
		currentGroup = append(currentGroup, elf)
		if len(currentGroup) == 3 {
			b := findBadge(currentGroup)
			s += priority(b)
			currentGroup = []string{}
		}
	}
	return s
}

func findDuplicateItem(contents string) rune {
	r := []rune(contents)
	r1 := r[:len(r)/2]
	r2 := r[len(r)/2:]
	for _, char := range r1 {
		if slices.Contains(r2, char) {
			return char
		}
	}
	return 0
}

func findBadge(group []string) rune {
	elf1 := []rune(group[0])
	elf2 := []rune(group[1])
	elf3 := []rune(group[2])
	for _, char := range elf1 {
		if slices.Contains(elf2, char) && slices.Contains(elf3, char) {
			return char
		}
	}
	return 0
}

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - 96
	} else {
		return int(r) - 38
	}
}
