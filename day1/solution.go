package day1

import (
    "sort"
    "strconv"

	"github.com/hpompecki/aoc_2022/input"
)

func MostCaloriesElf(inputFile string) int {
    allCalories := sumCalories(inputFile)
    return allCalories[0]
}

func Top3Calories(inputFile string) int {
    allCalories := sumCalories(inputFile)
    return allCalories[0] + allCalories[1] + allCalories[2]
}

func sumCalories(inputFile string) []int {
    allCalories := make([]int, 0)
    var currentCalories int
    for calories := range input.Read(inputFile) {
        if calories == "" {
            allCalories = append(allCalories, currentCalories)
            currentCalories = 0
            continue
        }
        c, _ := strconv.Atoi(calories)
        currentCalories += c
    }
    allCalories = append(allCalories, currentCalories)
    sort.Slice(allCalories, func(i, j int) bool {return allCalories[i] > allCalories[j]})
    return allCalories
}
