package main

import (
	"fmt"

	"github.com/hpompecki/aoc_2022/day1"
	"github.com/hpompecki/aoc_2022/day10"
	"github.com/hpompecki/aoc_2022/day11"
	"github.com/hpompecki/aoc_2022/day12"
	"github.com/hpompecki/aoc_2022/day13"
	"github.com/hpompecki/aoc_2022/day2"
	"github.com/hpompecki/aoc_2022/day3"
	"github.com/hpompecki/aoc_2022/day4"
	"github.com/hpompecki/aoc_2022/day5"
	"github.com/hpompecki/aoc_2022/day6"
	"github.com/hpompecki/aoc_2022/day7"
	"github.com/hpompecki/aoc_2022/day8"
	"github.com/hpompecki/aoc_2022/day9"
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

	day8Result1 := day8.CountVisibleTrees("day8/input.txt")
	fmt.Printf("Day 8, puzzle 1 result: %v\n", day8Result1)

	day8Result2 := day8.HighestScenicScore("day8/input.txt")
	fmt.Printf("Day 8, puzzle 2 result: %v\n", day8Result2)

	day9Result1 := day9.VisitedCount("day9/input.txt")
	fmt.Printf("Day 9, puzzle 1 result: %v\n", day9Result1)

	day9Result2 := day9.VisitedCountLong("day9/input.txt")
	fmt.Printf("Day 9, puzzle 2 result: %v\n", day9Result2)

	day10Result1 := day10.SumOfStrengths("day10/input.txt")
	fmt.Printf("Day 10, puzzle 1 result: %v\n", day10Result1)

	day10Result2 := day10.Render("day10/input.txt")
	fmt.Println("Day 10, puzzle 2 result:")
	for _, l := range day10Result2 {
		fmt.Println(string(l[:]))
	}

	day11Result1 := day11.MonkeyBusiness("day11/input.txt")
	fmt.Printf("Day 11, puzzle 1 result: %v\n", day11Result1)

	day11Result2 := day11.MonkeyBusiness2("day11/input.txt")
	fmt.Printf("Day 11, puzzle 2 result: %v\n", day11Result2)

	day12Result1 := day12.ShortestPath("day12/input.txt")
	fmt.Printf("Day 12, puzzle 1 result: %v\n", day12Result1)

	day12Result2 := day12.ShortestPathFromBestStart("day12/input.txt")
	fmt.Printf("Day 12, puzzle 2 result: %v\n", day12Result2)

	day13Result1 := day13.RightOrderIndices("day13/input.txt")
	fmt.Printf("Day 13, puzzle 1 result: %v\n", day13Result1)

	day13Result2 := day13.DecoderKey("day13/input.txt")
	fmt.Printf("Day 13, puzzle 2 result: %v\n", day13Result2)
}
