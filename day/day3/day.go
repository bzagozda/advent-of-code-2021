package day3

import (
	"github.com/bzagozda/aoc2021/util"
)

type Day3 struct{}

func (d *Day3) Number() int {
	return 3
}

func (d *Day3) Part1(input []string) (result int) {
	bytesArray := util.To2DByteArray(input)

	var gamma string
	var epsilon string

	for idx := range bytesArray[0] {
		bit, _ := mostCommonBit(bytesArray, idx)

		if bit == '0' {
			gamma += string('0')
			epsilon += string('1')
		} else {
			gamma += string('1')
			epsilon += string('0')
		}
	}

	return util.BinaryToInt(gamma) * util.BinaryToInt(epsilon)
}

func (d *Day3) Part2(input []string) (result int) {
	bytesArray := util.To2DByteArray(input)

	currentBytesArray := make([][]byte, len(bytesArray))
	copy(currentBytesArray, bytesArray)
	oxygenRating := getRating(currentBytesArray, '1', mostCommonBit)

	currentBytesArray = make([][]byte, len(bytesArray))
	copy(currentBytesArray, bytesArray)
	carbonRating := getRating(currentBytesArray, '0', leastCommonBit)

	return util.BinaryToInt(string(oxygenRating)) * util.BinaryToInt(string(carbonRating))
}

func mostCommonBit(bytes [][]byte, bitIdx int) (bit byte, half bool) {
	return getCommonBit(bytes, bitIdx, func(x, y int) bool { return x > y })
}

func leastCommonBit(bytes [][]byte, bitIdx int) (bit byte, half bool) {
	return getCommonBit(bytes, bitIdx, func(x, y int) bool { return x < y })
}

func getRating(bytesArray [][]byte, halfCriteria byte, criteriaFunction func([][]byte, int) (byte, bool)) []byte {
	cols := len(bytesArray[0])
	for idx := 0; idx < cols; idx++ {
		if len(bytesArray) == 1 {
			break
		}

		bit, half := criteriaFunction(bytesArray, idx)

		var carbonCriteria byte

		if half {
			carbonCriteria = halfCriteria
		} else {
			carbonCriteria = bit
		}

		bytesArray = filter(bytesArray, func(array []byte) bool {
			return array[idx] == carbonCriteria
		})
	}
	return bytesArray[0]
}

func getCommonBit(bytes [][]byte, bitIdx int, compFunc func(int, int) bool) (bit byte, half bool) {
	var zeros int
	var ones int

	for _, byteArray := range bytes {
		if byteArray[bitIdx] == '0' {
			zeros += 1
		} else {
			ones += 1
		}
	}

	half = zeros == ones

	if compFunc(zeros, ones) {
		bit = '0'
		return
	}

	bit = '1'
	return
}

func filter(bytes [][]byte, predicate func([]byte) bool) (result [][]byte) {
	for _, byteArray := range bytes {
		if predicate(byteArray) {
			result = append(result, byteArray)
		}
	}
	return
}
