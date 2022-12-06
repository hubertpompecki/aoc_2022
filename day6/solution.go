package day6

import (
	"github.com/hpompecki/aoc_2022/input"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/utf8string"
)

func PacketMarkerLocation(fileName string) int {
	line := utf8string.NewString(<-input.Read(fileName))
	for i := 3; i < line.RuneCount(); i++ {
		if uniqueLetters(line.Slice(i-3, i+1)) {
			return i + 1
		}
	}
	return 0
}

func MessageMarkerLocation(fileName string) int {
	line := utf8string.NewString(<-input.Read(fileName))
	for i := 13; i < line.RuneCount(); i++ {
		if uniqueLetters(line.Slice(i-13, i+1)) {
			return i + 1
		}
	}
	return 0
}

func uniqueLetters(s string) bool {
	seen := make([]rune, 0)
	for _, r := range s {
		if slices.Contains(seen, r) {
			return false
		}
		seen = append(seen, r)
	}
	return true
}
