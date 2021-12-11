package day4

import (
	"github.com/bzagozda/aoc2021/util"
)

type Day struct{}

func (d *Day) Number() int {
	return 4
}

func (d *Day) Part1(input []string) (result int) {
	bingoInput := parseInput(input)

	for _, draw := range bingoInput.choices {
		for _, board := range bingoInput.boards {
			for idx := range board {
				number := &(board)[idx]
				if *number == draw {
					*number = 0
				}

				if checkVictory(board, idx) {
					return getScore(board, draw)
				}
			}
		}
	}

	return
}

func (d *Day) Part2(input []string) (result int) {
	bingoInput := parseInput(input)

	boardWin := make([]bool, len(bingoInput.boards))

	for _, draw := range bingoInput.choices {
		for b := range bingoInput.boards {
			if boardWin[b] {
				continue
			}
			board := bingoInput.boards[b]
			for idx := range board {
				number := &(board)[idx]
				if *number == draw {
					*number = 0
				}

				if checkVictory(board, idx) {
					boardWin[b] = true
					winningCount := util.Count(len(boardWin), func(i int) bool {
						return boardWin[i]
					})

					if winningCount == len(bingoInput.boards) {
						return getScore(board, draw)
					}
				}
			}
		}
	}

	return
}

func getScore(board BingoBoard, multiplier int) (score int) {
	return util.ArraySum(board) * multiplier
}

func checkVictory(board BingoBoard, idx int) bool {
	row := idx / 5
	column := idx % 5
	rowData := board[row*5 : row*5+5]
	columnData := []int{
		board[column],
		board[column+5],
		board[column+10],
		board[column+15],
		board[column+20],
	}
	return util.ArraySum(rowData) == 0 || util.ArraySum(columnData) == 0
}
