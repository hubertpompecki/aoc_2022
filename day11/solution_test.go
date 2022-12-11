package day11

import (
	"testing"
)

func TestMonkeyBusiness(t *testing.T) {
	if r := MonkeyBusiness("input_test.txt"); r != 10605 {
		t.Fatalf("Expected 10605, got %v", r)
	}
}

func TestMonkeyBusiness2(t *testing.T) {
	if r := MonkeyBusiness2("input_test.txt"); r != 2713310158 {
		t.Fatalf("Expected 2713310158, got %v", r)
	}
}
