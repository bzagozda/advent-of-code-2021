package day7

import (
	"testing"

	"github.com/bzagozda/aoc2021/util"
)

func TestDay2Part1(t *testing.T) {
	input := []string{
		"16,1,2,0,4,2,7,1,2,14",
	}
	expected := 37

	day := &Day7{}
	result := day.Part1(input)
	util.Assert(t, result, expected)
}

func TestDay2Part2(t *testing.T) {
	input := []string{
		"16,1,2,0,4,2,7,1,2,14",
	}
	expected := 168

	day := &Day7{}
	result := day.Part2(input)
	util.Assert(t, result, expected)
}
