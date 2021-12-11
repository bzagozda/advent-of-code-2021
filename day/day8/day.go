package day8

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bzagozda/aoc2021/util"
)

type Day struct{}

func (d *Day) Number() int {
	return 8
}

func (d *Day) Part1(input []string) (count int) {
	for _, line := range input {
		output := strings.Split(strings.Trim(strings.Split(line, "|")[1], " "), " ")

		count += util.Count(len(output), func(i int) bool {
			segments := len(output[i])
			return segments == 2 || segments == 3 || segments == 4 || segments == 7
		})
	}
	return
}

func findDataByNoSegments(input []string) map[int][]string {
	uniqueSegmented := make(map[int][]string)
	for _, in := range input {
		if data, ok := uniqueSegmented[len(in)]; ok {
			uniqueSegmented[len(in)] = append(data, in)
		} else {
			uniqueSegmented[len(in)] = []string{in}
		}
	}
	return uniqueSegmented
}

func (d *Day) Part2(input []string) (count int) {
	for _, line := range input {
		inputOutput := strings.Split(line, "|")

		input := strings.Split(strings.Trim(inputOutput[0], " "), " ")
		output := strings.Split(strings.Trim(inputOutput[1], " "), " ")

		// segments := []string{".", ".", ".", ".", ".", ".", "."}

		dataByNoSegments := findDataByNoSegments(input)

		numbers := make(map[int]string, 0)
		numbers[1] = dataByNoSegments[2][0]
		delete(dataByNoSegments, 2)
		numbers[7] = dataByNoSegments[3][0]
		delete(dataByNoSegments, 3)
		numbers[4] = dataByNoSegments[4][0]
		delete(dataByNoSegments, 4)
		numbers[8] = dataByNoSegments[7][0]
		delete(dataByNoSegments, 7)
		if idx, err := findNumberIdxByNonCommonSegments(numbers[1], dataByNoSegments[6], 6); err == nil {
			numbers[6] = dataByNoSegments[6][idx]
			dataByNoSegments[6] = append(dataByNoSegments[6][:idx], dataByNoSegments[6][idx+1:]...)
		}
		if idx, err := findNumberIdxByNonCommonSegments(numbers[4], dataByNoSegments[6], 4); err == nil {
			numbers[0] = dataByNoSegments[6][idx]
			dataByNoSegments[6] = append(dataByNoSegments[6][:idx], dataByNoSegments[6][idx+1:]...)
		}
		numbers[9] = dataByNoSegments[6][0]
		delete(dataByNoSegments, 6)
		if idx, err := findNumberIdxByNonCommonSegments(numbers[1], dataByNoSegments[5], 3); err == nil {
			numbers[3] = dataByNoSegments[5][idx]
			dataByNoSegments[5] = append(dataByNoSegments[5][:idx], dataByNoSegments[5][idx+1:]...)
		}
		if idx, err := findNumberIdxByNonCommonSegments(numbers[6], dataByNoSegments[5], 1); err == nil {
			numbers[5] = dataByNoSegments[5][idx]
			dataByNoSegments[5] = append(dataByNoSegments[5][:idx], dataByNoSegments[5][idx+1:]...)
		}
		numbers[2] = dataByNoSegments[5][0]
		delete(dataByNoSegments, 5)

		//   0
		//  1 2
		//   3
		//  4 5
		//   6
		// segments[0] = nonCommonSegments(numbers[7], numbers[1])[0]
		// segments[2] = nonCommonSegments(numbers[5], numbers[9])[0]
		// segments[3] = nonCommonSegments(numbers[8], numbers[0])[0]
		// segments[4] = nonCommonSegments(numbers[8], numbers[9])[0]
		// segments[5] = nonCommonSegments(numbers[7], segments[0]+segments[2])[0]
		// segments[6] = nonCommonSegments(numbers[2], segments[1]+segments[2]+segments[3]+segments[4]+segments[6])[0]
		// segments[1] = nonCommonSegments(numbers[0], segments[0]+segments[2]+segments[4]+segments[5]+segments[6])[0]

		// drawSegments(segments)

		numberText := ""

		for _, out := range output {
			if num, err := findNumberIdxByNonCommonSegments(out, mapToList(numbers), 0); err == nil {
				numberText += fmt.Sprint(num)
			}
		}

		number, err := strconv.Atoi(numberText)
		if err != nil {
			log.Panic(err)
		}
		count += number
	}
	return
}

func mapToList(data map[int]string) (result []string) {
	for i := 0; i < len(data); i++ {
		result = append(result, data[i])
	}
	return
}

func drawSegments(segments []string) {
	log.Printf(" %s%s%s%s \n", segments[0], segments[0], segments[0], segments[0])
	log.Printf("%s    %s\n", segments[1], segments[2])
	log.Printf("%s    %s\n", segments[1], segments[2])
	log.Printf(" %s%s%s%s \n", segments[3], segments[3], segments[3], segments[3])
	log.Printf("%s    %s\n", segments[4], segments[5])
	log.Printf("%s    %s\n", segments[4], segments[5])
	log.Printf(" %s%s%s%s \n", segments[6], segments[6], segments[6], segments[6])
}

func nonCommonSegments(strings ...string) (result []string) {
	counts := make(map[rune]int)

	for _, str := range strings {
		for _, char := range str {
			counts[char] += 1
		}
	}

	for key, value := range counts {
		if value == 1 {
			result = append(result, string(key))
		}
	}

	return
}

func findNumberIdxByNonCommonSegments(number string, otherNumbers []string, requiredNonCommon int) (int, error) {
	for idx, num := range otherNumbers {
		nonCommonCharacters := nonCommonSegments(number, num)

		if len(nonCommonCharacters) == requiredNonCommon {
			return idx, nil
		}
	}

	return -1, fmt.Errorf("not found number with non common segments required %d", requiredNonCommon)
}
