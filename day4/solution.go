package day4

import (
	"github.com/hpompecki/aoc_2022/input"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func CountFullyOverlappingRanges(fileName string) int {
	var s int
	for l := range input.Read(fileName) {
		r1s, r1e, r2s, r2e := parseLine(l)
		if (r1s <= r2s && r1e >= r2e) || (r2s <= r1s && r2e >= r1e) {
			s++
		}
	}
	return s
}

func CountOverlappingRanges(fileName string) int {
	var s int
	for l := range input.Read(fileName) {
		r1s, r1e, r2s, r2e := parseLine(l)
		if (r1s <= r2s && r1e >= r2s) || (r2s <= r1s && r2e >= r1s) {
			s++
		}
	}
	return s
}

func parseLine(l string) (int, int, int, int) {
	m := re.FindSubmatch([]byte(l))
	r1s, _ := strconv.Atoi(string(m[1]))
	r1e, _ := strconv.Atoi(string(m[2]))
	r2s, _ := strconv.Atoi(string(m[3]))
	r2e, _ := strconv.Atoi(string(m[4]))
	return r1s, r1e, r2s, r2e
}
