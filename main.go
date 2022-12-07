package main

import (
	"fmt"

	"github.com/hpompecki/aoc_2022/day1"
	"github.com/hpompecki/aoc_2022/day2"
	"github.com/hpompecki/aoc_2022/day3"
	"github.com/hpompecki/aoc_2022/day4"
	"github.com/hpompecki/aoc_2022/day5"
	"github.com/hpompecki/aoc_2022/day6"
	"github.com/hpompecki/aoc_2022/day7"
)

func main() {
	day1Result1 := day1.MostCaloriesElf("day1/input.txt")
	fmt.Printf("Day 1, puzzle 1 result: %v\n", day1Result1)

	day1Result2 := day1.Top3Calories("day1/input.txt")
	fmt.Printf("Day 1, puzzle 2 result: %v\n", day1Result2)

	day2Result1 := day2.TotalPoints("day2/input.txt")
	fmt.Printf("Day 2, puzzle 1 result: %v\n", day2Result1)

	day2Result2 := day2.TotalPoints2("day2/input.txt")
	fmt.Printf("Day 2, puzzle 2 result: %v\n", day2Result2)

	day3Result1 := day3.SumOfPriorities("day3/input.txt")
	fmt.Printf("Day 3, puzzle 1 result: %v\n", day3Result1)

	day3Result2 := day3.SumOfBadgePriorities("day3/input.txt")
	fmt.Printf("Day 3, puzzle 2 result: %v\n", day3Result2)

	day4Result1 := day4.CountFullyOverlappingRanges("day4/input.txt")
	fmt.Printf("Day 4, puzzle 1 result: %v\n", day4Result1)

	day4Result2 := day4.CountOverlappingRanges("day4/input.txt")
	fmt.Printf("Day 4, puzzle 2 result: %v\n", day4Result2)

	day5Result1 := day5.TopCratesAfterMoves("day5/input.txt", true)
	fmt.Printf("Day 5, puzzle 1 result: %v\n", day5Result1)

	day5Result2 := day5.TopCratesAfterMoves("day5/input.txt", false)
	fmt.Printf("Day 5, puzzle 2 result: %v\n", day5Result2)

	day6Result1 := day6.PacketMarkerLocation("day6/input.txt")
	fmt.Printf("Day 6, puzzle 1 result: %v\n", day6Result1)

	day6Result2 := day6.MessageMarkerLocation("day6/input.txt")
	fmt.Printf("Day 6, puzzle 2 result: %v\n", day6Result2)

	day7Result1 := day7.SmallDirsTotal("day7/input.txt")
	fmt.Printf("Day 7, puzzle 1 result: %v\n", day7Result1)

	day7Result2 := day7.DirToDelete("day7/input.txt")
	fmt.Printf("Day 7, puzzle 2 result: %v\n", day7Result2)
}
