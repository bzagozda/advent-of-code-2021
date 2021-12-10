package day6

import (
	"strings"

	"github.com/bzagozda/aoc2021/util"
)

type Day6 struct{}

type MemKey struct {
	fish int
	days int
}

func (d *Day6) Number() int {
	return 6
}

func calculateFishes(fish int, days int, mem map[MemKey]int) int {
	if e, ok := mem[MemKey{days, fish}]; ok {
		return e
	}

	if days <= 0 || days <= fish {
		mem[MemKey{days, fish}] = 1
		return 1
	}

	daysAfterBirth := (days - fish - 1)

	fishes := calculateFishes(6, daysAfterBirth, mem) + calculateFishes(8, daysAfterBirth, mem)

	mem[MemKey{days, fish}] = fishes
	return fishes
}

func (d *Day6) Part1(input []string) (result int) {
	initialState := util.MapToInt(strings.Split(input[0], ","))

	mem := make(map[MemKey]int, 0)

	fishes := util.ArraySum(util.ApplyInt(initialState, func(fish int) int {
		return calculateFishes(fish, 80, mem)
	}))

	return fishes
}

func (d *Day6) Part2(input []string) (result int) {
	initialState := util.MapToInt(strings.Split(input[0], ","))

	mem := make(map[MemKey]int, 0)

	fishes := util.ArraySum(util.ApplyInt(initialState, func(fish int) int {
		return calculateFishes(fish, 256, mem)
	}))

	return fishes
}
