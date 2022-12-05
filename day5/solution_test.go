package day5

import (
	"testing"
)

func TestMovingCratesBy1(t *testing.T) {
	if s := TopCratesAfterMoves("input_test.txt", true); s != "CMZ" {
		t.Fatalf("Expected CMZ, got: %v", s)
	}
}

func TestMovingCratesInBulk(t *testing.T) {
	if s := TopCratesAfterMoves("input_test.txt", false); s != "MCD" {
		t.Fatalf("Expected MCD, got: %v", s)
	}
}
