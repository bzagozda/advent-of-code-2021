package util

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func SlidingWindow(data []int, window int) (result [][]int) {
	if window < 1 {
		return result
	}

	for i := 0; i < len(data)-window+1; i += 1 {
		windowData := make([]int, window)
		for x := 0; x < window; x += 1 {
			windowData[x] = data[i+x]
		}
		result = append(result, windowData)
	}

	return
}

func Count(times int, predFunc func(int) bool) (count int) {
	for i := 0; i < times; i++ {
		if predFunc(i) {
			count += 1
		}
	}
	return
}

func Count2D(maxX, maxY int, predFunc func(x, y int) bool) (count int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if predFunc(x, y) {
				count += 1
			}
		}
	}
	return
}

func MapToInt(input []string) []int {
	data := make([]int, 0)

	for _, rawNumber := range input {
		if number, err := strconv.Atoi(rawNumber); err == nil {
			data = append(data, number)
		}
	}

	return data
}

func ApplyInt(input []int, applyFunc func(int) int) []int {
	data := make([]int, 0)

	for _, number := range input {
		data = append(data, applyFunc(number))
	}

	return data
}

func ApplyString(input []string, applyFunc func(string) string) []string {
	data := make([]string, 0)

	for _, number := range input {
		data = append(data, applyFunc(number))
	}

	return data
}

func ArraySum(input []int) (sum int) {
	for _, num := range input {
		sum += num
	}
	return
}

func To2DByteArray(input []string) (result [][]byte) {
	for _, line := range input {
		result = append(result, []byte(line))
	}
	return
}

func To2DRuneArray(input []string) (result [][]rune) {
	for _, line := range input {
		result = append(result, []rune(line))
	}
	return
}

func To2DIntArray(input []string) (result [][]int) {
	for _, line := range input {
		row := make([]int, 0)
		for _, e := range line {
			num := int(e - '0')
			row = append(row, num)
		}
		result = append(result, row)
	}
	return
}

func Print2DByteArray(arr [][]byte) {
	for _, row := range arr {
		log.Println(string(row))
	}
}

func Print2DRuneArray(arr [][]rune) {
	for _, row := range arr {
		log.Println(string(row))
	}
}

func Print2DIntArray(arr [][]int) {
	for _, row := range arr {
		log.Println(row)
	}
}

func StringArrayToRunes(arr []string) (runes []rune) {
	for _, str := range arr {
		for _, ch := range str {
			runes = append(runes, ch)
		}
	}
	return
}

func Iter2D(arr [][]byte, iterFunc func(y, x int, el byte)) {
	for y, row := range arr {
		for x, el := range row {
			iterFunc(y, x, el)
		}
	}
}

func BinaryToInt(binary string) int {
	num, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Println(err)
	}

	return int(num)
}

func StringToIntArray(textArray string) (result []int) {
	for _, n := range strings.Split(textArray, ",") {
		number, err := strconv.Atoi(n)
		if err != nil {
			continue
		}
		result = append(result, number)
	}

	return
}

func MinAndMax(array []int) (min, max int) {
	min = math.MaxInt
	max = math.MinInt

	for _, n := range array {
		min = int(math.Min(float64(n), float64(min)))
		max = int(math.Max(float64(n), float64(max)))
	}

	return
}

func ForEach(times int, iterFunc func(i int)) {
	for i := 0; i < times; i++ {
		iterFunc(i)
	}
}
