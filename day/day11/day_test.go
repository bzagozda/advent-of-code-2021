package day11

import (
	"testing"

	"github.com/bzagozda/aoc2021/util"
)

func TestDay2Part1(t *testing.T) {
	input := []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}
	expected := 1656

	day := &Day{}
	result := day.Part1(input)
	util.Assert(t, result, expected)
}

func TestDay2Part2(t *testing.T) {
	input := []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}
	expected := 195

	day := &Day{}
	result := day.Part2(input)
	util.Assert(t, result, expected)
}
