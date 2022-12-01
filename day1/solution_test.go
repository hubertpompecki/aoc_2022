package day1

import (
    "testing"
)

func TestMostCalories(t *testing.T) {
    if mostCalories := MostCaloriesElf("input_test.txt"); mostCalories != 24000 {
       t.Fatalf("Expected 24000, got: %v", mostCalories)
   }
}

func TestSumOfTop3(t *testing.T) {
    if top3Calories := Top3Calories("input_test.txt"); top3Calories != 45000 {
       t.Fatalf("Expected 45000, got: %v", top3Calories)
   }
}
