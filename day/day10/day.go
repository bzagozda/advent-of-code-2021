package day10

import (
	"fmt"
	"log"
	"sort"
)

type Day struct{}

var (
	scoring map[rune]int = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	stackScoring map[rune]int = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	chunkChars map[rune]rune = map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
		'<': '>',
	}
)

type Stack struct {
	stack []rune
}

func (s *Stack) Push(char rune) {
	s.stack = append(s.stack, char)
}

func (s *Stack) Peek() (rune, error) {
	if len(s.stack) == 0 {
		return ' ', fmt.Errorf("Stack is empty")
	}
	return s.stack[len(s.stack)-1], nil
}

func (s *Stack) Pop() (rune, error) {
	if len(s.stack) == 0 {
		return ' ', fmt.Errorf("Stack is empty")
	}

	char := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]

	return char, nil
}

func (d *Day) Number() int {
	return 10
}

func findIllegalChar(text []rune) (char rune, found bool) {
	stack := Stack{}

	for _, char := range text {
		if _, ok := chunkChars[char]; !ok {
			opening, err := stack.Peek()

			if err != nil {
				log.Panic(err)
			}

			if correspondingChar, ok := chunkChars[opening]; ok && correspondingChar == char {
				_, err := stack.Pop()
				if err != nil {
					log.Panic(err)
				}
			} else {
				return char, true
			}
		} else {
			stack.Push(char)
		}
	}

	return '0', false
}

func calculateScore(text []rune) (scoreSum int) {
	for _, char := range text {
		if score, ok := scoring[char]; ok {
			scoreSum += score
		}
	}
	return
}

func (d *Day) Part1(input []string) (result int) {
	illegalCharacters := make([]rune, 0)
	for _, line := range input {
		if char, found := findIllegalChar([]rune(line)); found {
			illegalCharacters = append(illegalCharacters, char)
		}
	}
	return calculateScore(illegalCharacters)
}

func findMissingCompletions(text []rune) *Stack {
	stack := Stack{}

	for _, char := range text {
		if _, ok := chunkChars[char]; !ok {
			opening, err := stack.Peek()

			if err != nil {
				log.Panic(err)
			}

			if correspondingChar, ok := chunkChars[opening]; ok && correspondingChar == char {
				_, err := stack.Pop()
				if err != nil {
					log.Panic(err)
				}
			}
		} else {
			stack.Push(char)
		}
	}

	return &stack
}

func calculateStackScore(s *Stack) (scoreSum int) {
	for char, err := s.Pop(); err == nil; char, err = s.Pop() {
		if correspondingChar, ok := chunkChars[char]; ok {
			if score, ok := stackScoring[correspondingChar]; ok {
				scoreSum = scoreSum*5 + score
			}
		}
	}
	return
}

func (d *Day) Part2(input []string) (result int) {
	scores := make([]int, 0)
	for _, line := range input {
		if _, found := findIllegalChar([]rune(line)); !found {
			s := findMissingCompletions([]rune(line))
			scores = append(scores, calculateStackScore(s))
		}
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}
