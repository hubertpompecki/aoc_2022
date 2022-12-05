package day5

import (
	"regexp"
	"strconv"

	"github.com/hpompecki/aoc_2022/input"
	"golang.org/x/exp/utf8string"
)

var re = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func TopCratesAfterMoves(fileName string, reorder bool) string {
	var stacks = make([][]rune, 0)

	// parse input
	lines := input.Read(fileName)
	for l := range lines {
		// build new stacks when parsing the first line
		if len(stacks) == 0 {
			for i := 0; i <= len(l)/4; i++ {
				stacks = append(stacks, make([]rune, 0))
			}
		}

		// make the string easier to work with
		l := utf8string.NewString(l)

		// check if we finished bulding stacks
		if l.At(1) == '1' {
			break
		}

		putLineOnStacks(stacks, l)
	}

	// skip blank line before moving instructions start
	<-lines

	// execute moving instructions
	for l := range lines {
		moveCrates(stacks, l, reorder)
	}

	result := make([]rune, 0)
	for _, stack := range stacks {
		result = append(result, stack[0])
	}

	return string(result)
}

func putLineOnStacks(stacks [][]rune, l *utf8string.String) {
	for i := range stacks {
		r := l.At(i*4 + 1)
		if r != ' ' {
			stacks[i] = append(stacks[i], r)
		}
	}
}

func moveCrates(stacks [][]rune, l string, reorder bool) {
	m := re.FindSubmatch([]byte(l))
	n, _ := strconv.Atoi(string(m[1]))
	from, _ := strconv.Atoi(string(m[2]))
	to, _ := strconv.Atoi(string(m[3]))

	moved := append([]rune{}, stacks[from-1][:n]...)
	stacks[from-1] = stacks[from-1][n:]
	if reorder {
		reverse(moved)
	}
	stacks[to-1] = append(moved, stacks[to-1]...)
}

func reverse(s []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
