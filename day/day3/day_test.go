package day3

import "testing"

func TestDay2Part1(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	expected := 198

	day := &Day3{}
	result := day.Part1(input)
	if result != expected {
		t.Errorf("Expected: %d, Got: %d", expected, result)
	}
}

func TestDay2Part2(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	expected := 230

	day := &Day3{}
	result := day.Part2(input)
	if result != expected {
		t.Errorf("Expected: %d, Got: %d", expected, result)
	}
}
