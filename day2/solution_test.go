package day2

import (
    "testing"
)

func TestTotalPoints(t *testing.T) {
    expectedPoints := 15
    if totalPoints := TotalPoints("input_test.txt"); totalPoints != expectedPoints {
       t.Fatalf("Expected %v, got: %v", expectedPoints, totalPoints)
   }
}
