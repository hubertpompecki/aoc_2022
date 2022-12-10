package day10

import (
	"fmt"
	"testing"
)

func TestSumOfStrengths(t *testing.T) {
	if s := SumOfStrengths("input_test.txt"); s != 13140 {
		t.Fatalf("Expected 13140, got %v", s)
	}
}

func TestRender(t *testing.T) {
	for _, l := range Render("input_test.txt") {
		fmt.Println(string(l[:]))
	}
}
