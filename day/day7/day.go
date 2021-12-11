package day7

import (
	"math"

	"github.com/bzagozda/aoc2021/util"
)

type Day struct{}

func (d *Day) Number() int {
	return 7
}

func (d *Day) Part1(input []string) int {
	positions := util.StringToIntArray(input[0])
	fuelUsageCalcFunc := func(pos1, pos2 int) int {
		return int(math.Abs(float64(pos1) - float64(pos2)))
	}
	return getMinFuel(positions, fuelUsageCalcFunc)
}

func (d *Day) Part2(input []string) int {
	positions := util.StringToIntArray(input[0])
	fuelUsageCalcFunc := func(pos1, pos2 int) int {
		diff := math.Abs(float64(pos1) - float64(pos2))
		return int(diff * ((1.0 + diff) / 2.0))
	}
	return getMinFuel(positions, fuelUsageCalcFunc)
}

func getMinFuel(positions []int, fuelUsageCalcFunc func(pos1, pos2 int) int) int {
	min, max := util.MinAndMax(positions)

	minFuel := math.MaxInt

	for i := min; i <= max; i++ {
		fuelSum := 0
		for _, pos := range positions {
			fuelSum += fuelUsageCalcFunc(i, pos)
		}
		minFuel = int(math.Min(float64(minFuel), float64(fuelSum)))
	}

	return minFuel
}
