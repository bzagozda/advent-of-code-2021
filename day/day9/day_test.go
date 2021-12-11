package day9

import (
	"testing"

	"github.com/bzagozda/aoc2021/util"
)

func TestDay2Part1(t *testing.T) {
	input := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}
	expected := 15

	day := &Day{}
	result := day.Part1(input)
	util.Assert(t, result, expected)
}

func TestDay2Part2(t *testing.T) {
	input := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}
	expected := 1134

	day := &Day{}
	result := day.Part2(input)
	util.Assert(t, result, expected)
}
