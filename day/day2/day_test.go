package day2

import "testing"

func TestDay2Part1(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	expected := 150

	day2 := &Day2{}
	result := day2.Part1(input)
	if result != expected {
		t.Errorf("Expected: %d, Got: %d", expected, result)
	}
}

func TestDay2Part2(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	expected := 900

	day2 := &Day2{}
	result := day2.Part2(input)
	if result != expected {
		t.Errorf("Expected: %d, Got: %d", expected, result)
	}
}
