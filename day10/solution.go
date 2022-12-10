package day10

import (
	"math"
	"strconv"
	"strings"

	"github.com/hpompecki/aoc_2022/input"
)

func SumOfStrengths(fileName string) int {
	var (
		c = 1
		v = 1
		s int
	)
	for l := range input.Read(fileName) {
		switch {
		case l == "noop":
			checkStrength(c, v, &s)
			c++
			continue
		case strings.HasPrefix(l, "addx "):
			n, _ := strconv.Atoi(strings.Split(l, " ")[1])
			checkStrength(c, v, &s)
			c++
			checkStrength(c, v, &s)
			c++
			v = v + n
		}
	}
	return s
}

func Render(fileName string) [6][40]rune {
	screen := [6][40]rune{}
	var (
		c = 1
		v = 1
	)
	for l := range input.Read(fileName) {
		switch {
		case l == "noop":
			draw(&screen, v, c)
			c++
			continue
		case strings.HasPrefix(l, "addx "):
			n, _ := strconv.Atoi(strings.Split(l, " ")[1])
			draw(&screen, v, c)
			c++
			draw(&screen, v, c)
			c++
			v = v + n
		}
	}
	return screen
}

func draw(screen *[6][40]rune, v int, c int) {
	row := (c - 1) / 40
	position := (c - 1) % 40
	if (math.Abs(float64((v+1)-(c%40))) <= 1 && c%40 != 0) || (c%40 == 0 && (v == 38 || v == 39)) {
		screen[row][position] = '#'
	} else {
		screen[row][position] = '.'
	}
}

func checkStrength(c, v int, s *int) {
	if c == 20 || c == 60 || c == 100 || c == 140 || c == 180 || c == 220 {
		*s = *s + v*c
	}
}
