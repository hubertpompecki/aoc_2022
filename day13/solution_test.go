package day13

import "testing"

func TestRightOrderIndices(t *testing.T) {
    if r := RightOrderIndices("input_test.txt"); r != 13 {
        t.Fatalf("Expected 13, got :%v", r)
    }
}

func TestDecoderKey(t *testing.T) {
    if r := DecoderKey("input_test.txt"); r != 140 {
        t.Fatalf("Expected 140, got :%v", r)
    }
}
