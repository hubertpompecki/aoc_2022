package day7

import (
	"testing"
)

func TestFindingSmallDirs(t *testing.T) {
	if l := SmallDirsTotal("input_test.txt"); l != 95437 {
		t.Fatalf("1) Expected 795437, got: %v", l)
	}
}

func TestDirToDelete(t *testing.T) {
	if l := DirToDelete("input_test.txt"); l != 24933642 {
		t.Fatalf("1) Expected 24933642, got: %v", l)
	}
}
