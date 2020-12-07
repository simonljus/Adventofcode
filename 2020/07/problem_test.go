package main

import (
	"testing"
)

func TestToBag(t *testing.T) {
	color, bags := toBag("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	if color != "light red" {
		t.Errorf("expected [%s] was [%s]", "light red", color)
	}
	if len(bags) != 2 {
		t.Errorf("expected [%d] was [%d]", 2, len(bags))
	}
}
func TestSolve(t *testing.T) {

	p1Expected, p2Expected := int64(4), int64(32)
	if p1Actual, p2Actual := solve("test_input.txt"); p1Actual != p1Expected {
		t.Errorf("Problem1: expected  %d actual %d ", p1Expected, p1Actual)
	} else if p2Actual != p2Expected {
		t.Errorf("Problem2: expected  %d actual %d ", p2Expected, p2Actual)
	}
}
