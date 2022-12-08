package day8

import (
	"testing"
)

func TestCountingVisibleTrees(t *testing.T) {
	if n := CountVisibleTrees("input_test.txt"); n != 21 {
		t.Fatalf("Expected 21, got %v", n)
	}
}

func TestHighestScenicScore(t *testing.T) {
	if n := HighestScenicScore("input_test.txt"); n != 8 {
		t.Fatalf("Expected 8, got %v", n)
	}
}
