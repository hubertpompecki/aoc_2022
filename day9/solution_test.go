package day9

import (
    "testing"
)

func TestCountingVisitedPositions(t *testing.T) {
    if r := VisitedCount("input_test.txt"); r != 13 {
        t.Fatalf("Expected 13, got %v", r)
    }
}
func TestCountingVisitedPositionsLong(t *testing.T) {
    if r := VisitedCountLong("input_test.txt"); r != 1 {
        t.Fatalf("Expected 1, got %v", r)
    }

    if r := VisitedCountLong("input_test_2.txt"); r != 36 {
        t.Fatalf("Expected 36, got %v", r)
    }
}
