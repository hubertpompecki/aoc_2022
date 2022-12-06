package day6

import (
	"testing"
)

func TestFindingPacketMarker(t *testing.T) {
	if l := PacketMarkerLocation("input_test_1.txt"); l != 7 {
		t.Fatalf("1) Expected 7, got: %v", l)
	}

	if l := PacketMarkerLocation("input_test_2.txt"); l != 5 {
		t.Fatalf("2) Expected 5, got: %v", l)
	}

	if l := PacketMarkerLocation("input_test_3.txt"); l != 6 {
		t.Fatalf("3) Expected 6, got: %v", l)
	}

	if l := PacketMarkerLocation("input_test_4.txt"); l != 10 {
		t.Fatalf("4) Expected 10, got: %v", l)
	}

	if l := PacketMarkerLocation("input_test_5.txt"); l != 11 {
		t.Fatalf("5) Expected 11, got: %v", l)
	}
}

func TestFindingMessageMarker(t *testing.T) {
	if l := MessageMarkerLocation("input_test_1.txt"); l != 19 {
		t.Fatalf("1) Expected 19, got: %v", l)
	}

	if l := MessageMarkerLocation("input_test_2.txt"); l != 23 {
		t.Fatalf("2) Expected 23, got: %v", l)
	}

	if l := MessageMarkerLocation("input_test_3.txt"); l != 23 {
		t.Fatalf("3) Expected 23, got: %v", l)
	}

	if l := MessageMarkerLocation("input_test_4.txt"); l != 29 {
		t.Fatalf("4) Expected 29, got: %v", l)
	}

	if l := MessageMarkerLocation("input_test_5.txt"); l != 26 {
		t.Fatalf("5) Expected 26, got: %v", l)
	}
}
