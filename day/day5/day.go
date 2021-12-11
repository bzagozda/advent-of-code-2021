package day5

import (
	"math"
	"sort"
	"strings"

	"github.com/bzagozda/aoc2021/util"
)

type Day struct{}

func (d *Day) Number() int {
	return 5
}

type Point struct {
	x int
	y int
}

type Vent struct {
	from Point
	to   Point
}

func parseInput(input []string) (vents []Vent) {
	for _, line := range input {
		points := strings.Split(line, " -> ")
		pointFrom := util.MapToInt(strings.Split(points[0], ","))
		pointTo := util.MapToInt(strings.Split(points[1], ","))
		vent := Vent{
			from: Point{
				x: pointFrom[0],
				y: pointFrom[1],
			},
			to: Point{
				x: pointTo[0],
				y: pointTo[1],
			},
		}
		vents = append(vents, vent)
	}

	return
}

func getMapSize(vents []Vent) (maxX, maxY int) {
	for _, vent := range vents {
		maxX = int(math.Max(float64(vent.to.x), math.Max(float64(vent.from.x), float64(maxX))))
		maxY = int(math.Max(float64(vent.to.y), math.Max(float64(vent.from.y), float64(maxY))))
	}
	maxX += 1
	maxY += 1
	return
}

func (d *Day) Part1(input []string) (result int) {
	vents := parseInput(input)
	maxX, maxY := getMapSize(vents)

	overlapMap := make([][]int, maxY)
	for i := range overlapMap {
		overlapMap[i] = make([]int, maxX)
	}

	for _, vent := range vents {
		if vent.from.x == vent.to.x {
			ys := []int{vent.from.y, vent.to.y}
			sort.Ints(ys)
			for i := 0; i <= ys[1]-ys[0]; i++ {
				overlapMap[ys[0]+i][vent.from.x]++
			}
		} else if vent.from.y == vent.to.y {
			xs := []int{vent.from.x, vent.to.x}
			sort.Ints(xs)
			for i := 0; i <= xs[1]-xs[0]; i++ {
				overlapMap[vent.from.y][xs[0]+i]++
			}
		}
	}

	// for _, line := range overlapMap {
	// 	log.Println(line)
	// }

	dangerousSpaces := util.Count2D(maxX, maxY, func(x, y int) bool {
		return overlapMap[y][x] >= 2
	})

	return dangerousSpaces
}

func (d *Day) Part2(input []string) (result int) {
	vents := parseInput(input)
	maxX, maxY := getMapSize(vents)

	overlapMap := make([][]int, maxY)
	for i := range overlapMap {
		overlapMap[i] = make([]int, maxX)
	}

	for _, vent := range vents {
		if vent.from.x == vent.to.x {
			ys := []int{vent.from.y, vent.to.y}
			sort.Ints(ys)
			for i := 0; i <= ys[1]-ys[0]; i++ {
				overlapMap[ys[0]+i][vent.from.x]++
			}
		} else if vent.from.y == vent.to.y {
			xs := []int{vent.from.x, vent.to.x}
			sort.Ints(xs)
			for i := 0; i <= xs[1]-xs[0]; i++ {
				overlapMap[vent.from.y][xs[0]+i]++
			}
		} else {
			lenX := int(math.Abs(float64(vent.to.x) - float64(vent.from.x)))
			lenY := int(math.Abs(float64(vent.to.y) - float64(vent.from.y)))
			signX := (vent.to.x - vent.from.x) / lenX
			signY := (vent.to.y - vent.from.y) / lenY

			for i := 0; i <= lenX; i++ {
				overlapMap[vent.from.y+(signY*i)][vent.from.x+(signX*i)]++
			}
		}
	}

	dangerousSpaces := util.Count2D(maxX, maxY, func(x, y int) bool {
		return overlapMap[y][x] >= 2
	})

	return dangerousSpaces
}
