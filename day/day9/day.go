package day9

import (
	"log"
	"math"
	"sort"
	"strconv"

	"github.com/bzagozda/aoc2021/util"
)

type Day struct{}

func (d *Day) Number() int {
	return 9
}

func (d *Day) Part1(input []string) (result int) {
	heightMap := util.To2DByteArray(input)
	util.Iter2D(heightMap, func(y, x int, el byte) {
		neighbors, _ := getNeighbors(heightMap, y, x)

		lowestPoint := el
		for _, n := range neighbors {
			if n == el {
				return
			}
			lowestPoint = byte(math.Min(float64(lowestPoint), float64(n)))
		}

		if lowestPoint == el {
			num, err := strconv.Atoi(string(el))
			if err != nil {
				log.Panic(err)
			}
			result += num + 1
		}
	})

	return
}

func getBasinSize(heightMap *[][]byte, y, x int) (basinSize int) {
	if (*heightMap)[y][x] == byte('9') {
		return
	}

	basinSize += 1
	(*heightMap)[y][x] = byte('9')

	_, indexes := getNeighbors(*heightMap, y, x)
	for idx := range indexes {
		basinSize += getBasinSize(heightMap, indexes[idx][0], indexes[idx][1])
	}

	return
}

func (d *Day) Part2(input []string) int {
	heightMap := util.To2DByteArray(input)
	basins := make([]int, 0)

	for y := range heightMap {
		for x, el := range heightMap[y] {
			if el == byte('9') {
				continue
			}

			basins = append(basins, getBasinSize(&heightMap, y, x))
		}
	}

	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	return basins[0] * basins[1] * basins[2]
}

func getNeighbors(heightmap [][]byte, y, x int) (neighbors []byte, indexes [][]int) {
	neighborIndexes := [][]int{
		{y - 1, x},
		{y, x - 1},
		{y + 1, x},
		{y, x + 1},
	}

	for _, ni := range neighborIndexes {
		if ni[0] >= 0 && ni[0] < len(heightmap) {
			if ni[1] >= 0 && ni[1] < len(heightmap[ni[0]]) {
				neighbors = append(neighbors, heightmap[ni[0]][ni[1]])
				indexes = append(indexes, ni)
			}
		}
	}

	return
}
