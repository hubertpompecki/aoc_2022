package day12

import (
	"testing"
)

func TestShortestPath(t *testing.T) {
	if r := ShortestPath("input_test.txt"); r != 31 {
		t.Fatalf("Expected 31, got %v", r)
	}
}

func TestShortestPathFromBestStart(t *testing.T) {
	if r := ShortestPathFromBestStart("input_test.txt"); r != 29 {
		t.Fatalf("Expected 29, got %v", r)
	}
}
