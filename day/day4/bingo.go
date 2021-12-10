package day4

import (
	"strings"

	"github.com/bzagozda/aoc2021/util"
)

type BingoBoard []int

type Choices []int

type BingoInput struct {
	choices Choices
	boards  []BingoBoard
}

func parseBoard(rawBoard []string) (bingoBoard BingoBoard) {
	for _, line := range rawBoard {
		numbers := util.MapToInt(strings.Split(line, " "))
		bingoBoard = append(bingoBoard, numbers...)
	}
	return
}

func parseInput(input []string) BingoInput {
	rawChoices := strings.Split(input[0], ",")
	choices := util.MapToInt(rawChoices)

	bingoBoards := make([]BingoBoard, 0)
	for i := 2; i < len(input); i += 6 {
		rawBoard := input[i : i+5]
		bingoBoard := parseBoard(rawBoard)
		bingoBoards = append(bingoBoards, bingoBoard)
	}

	return BingoInput{
		choices: choices,
		boards:  bingoBoards,
	}
}
