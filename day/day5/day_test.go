package day5

import "testing"

func TestDay2Part1(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	expected := 5

	day := &Day5{}
	result := day.Part1(input)
	if result != expected {
		t.Errorf("Expected: %d, Got: %d", expected, result)
	}
}

func TestDay2Part2(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	expected := 12

	day := &Day5{}
	result := day.Part2(input)
	if result != expected {
		t.Errorf("Expected: %d, Got: %d", expected, result)
	}
}
