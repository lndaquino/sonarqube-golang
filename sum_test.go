package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)
	if result != 5 {
		t.Error("Result must be 5")
	}
}

func TestSub(t *testing.T) {
	result := sub(3, 1)
	if result != 2 {
		t.Error("Result must be 2")
	}
}

func TestTimes(t *testing.T) {
	result := times(2, 3)
	if result != 6 {
		t.Error("Result must be 6")
	}
}

func TestSumX(t *testing.T) {
	result := sumX(2)
	if result != 4 {
		t.Error("Result must be 4")
	}
}
