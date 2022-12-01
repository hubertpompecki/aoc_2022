package main

import (
    "fmt"

	"github.com/hpompecki/aoc_2022/day1"
)

func main() {
    day1Result1 := day1.MostCaloriesElf("day1/input.txt")
    fmt.Printf("Day 1, puzzle 1 result: %v\n", day1Result1)

    day1Result2 := day1.Top3Calories("day1/input.txt")
    fmt.Printf("Day 1, puzzle 2 result: %v\n", day1Result2)
}
