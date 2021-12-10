package day6

import (
	"testing"

	"github.com/bzagozda/aoc2021/util"
)

func TestDay2Part1(t *testing.T) {
	input := []string{
		"3,4,3,1,2",
	}
	expected := 5934

	day := &Day6{}
	result := day.Part1(input)
	util.Assert(t, result, expected)
}

func TestDay2Part2(t *testing.T) {
	input := []string{
		"3,4,3,1,2",
	}
	expected := 26984457539

	day := &Day6{}
	result := day.Part2(input)
	util.Assert(t, result, expected)
}

func BenchmarkDay2Part2(b *testing.B) {
	input := []string{
		"3,4,3,1,2",
	}

	day := &Day6{}

	for n := 0; n < b.N; n++ {
		day.Part2(input)
	}
}
