package day4

import (
	"testing"
)

func TestCountFullyOverlappingRanges(t *testing.T) {
	if s := CountFullyOverlappingRanges("input_test.txt"); s != 2 {
		t.Fatalf("Expected 2, got: %v", s)
	}
}

func TestCountOverlappingRanges(t *testing.T) {
	if s := CountOverlappingRanges("input_test.txt"); s != 4 {
		t.Fatalf("Expected 4, got: %v", s)
	}
}
