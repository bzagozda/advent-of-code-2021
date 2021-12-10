package day1

import "github.com/bzagozda/aoc2021/util"

type Day1 struct{}

func (d *Day1) Number() int {
	return 1
}

func countIncreasing(meters []int) int {
	windows := util.SlidingWindow(meters, 2)
	return util.Count(len(windows), func(i int) bool { return windows[i][0] < windows[i][1] })
}

func (d *Day1) Part1(input []string) int {
	meters := util.MapToInt(input)
	return countIncreasing(meters)
}

func (d *Day1) Part2(input []string) int {
	meters := util.MapToInt(input)
	windows := util.SlidingWindow(meters, 3)
	summed := make([]int, len(windows))
	for i := range summed {
		summed[i] = util.ArraySum(windows[i])
	}

	return countIncreasing(summed)
}
