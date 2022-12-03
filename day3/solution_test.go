package day3

import (
	"testing"
)

func TestSumOfPriorities(t *testing.T) {
	if s := SumOfPriorities("input_test.txt"); s != 157 {
		t.Fatalf("Expected 157, got: %v", s)
	}
}

func TestSumOfBadgePriorities(t *testing.T) {
	if s := SumOfBadgePriorities("input_test.txt"); s != 70 {
		t.Fatalf("Expected 70, got: %v", s)
	}
}
