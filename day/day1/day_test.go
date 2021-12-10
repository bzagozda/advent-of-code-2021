package day1

import (
	"testing"
)

func TestDay1Part1(t *testing.T) {
	input := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}
	expected := 7

	day1 := &Day1{}
	result := day1.Part1(input)
	if result != expected {
		t.Errorf("Expected: %d, Got: %d", expected, result)
	}
}

func TestDay1Part2(t *testing.T) {
	input := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}
	expected := 5

	day1 := &Day1{}
	result := day1.Part2(input)
	if result != expected {
		t.Errorf("Expected: %d, Got: %d", expected, result)
	}
}
