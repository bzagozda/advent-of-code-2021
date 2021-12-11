package day11

import (
	"github.com/bzagozda/aoc2021/util"
)

type Day struct{}

func (d *Day) Number() int {
	return 11
}

func step(cave [][]int) (flashes int, allFlashes bool) {
	util.ForEach(len(cave), func(y int) {
		util.ForEach(len(cave[y]), func(x int) {
			cave[y][x] += 1

			if cave[y][x] == 10 {
				flashes += flash(cave, x, y)
			}
		})
	})

	if flashes == len(cave)*len(cave[0]) {
		allFlashes = true
	}

	util.ForEach(len(cave), func(y int) {
		util.ForEach(len(cave[y]), func(x int) {
			if cave[y][x] > 9 {
				cave[y][x] = 0
			}
		})
	})
	return
}

func flash(cave [][]int, x, y int) (flashes int) {
	flashes += 1

	applyToNeighbors(cave, x, y, func(neighborY, neighborX int) {
		neighbor := &cave[neighborY][neighborX]
		(*neighbor) += 1
		if (*neighbor) == 10 {
			flashes += flash(cave, neighborX, neighborY)
		}
	})
	return
}

func applyToNeighbors(cave [][]int, x, y int, applyFunc func(neighborY, neighborX int)) {
	neighbors := [][]int{
		{y - 1, x - 1},
		{y - 1, x},
		{y - 1, x + 1},
		{y + 0, x - 1},
		{y + 0, x + 1},
		{y + 1, x - 1},
		{y + 1, x},
		{y + 1, x + 1},
	}

	for _, neighbor := range neighbors {
		if neighbor[0] >= 0 && neighbor[0] < len(cave) && neighbor[1] >= 0 && neighbor[1] < len(cave[0]) {
			applyFunc(neighbor[0], neighbor[1])
		}
	}
}

func (d *Day) Part1(input []string) (result int) {
	cave := util.To2DIntArray(input)

	for i := 0; i < 100; i++ {
		flashes, _ := step(cave)
		result += flashes
	}

	return
}

func (d *Day) Part2(input []string) (result int) {
	cave := util.To2DIntArray(input)

	stepCount := 0
	for {
		stepCount++
		_, allFlashed := step(cave)
		if allFlashed {
			return stepCount
		}
	}
}
