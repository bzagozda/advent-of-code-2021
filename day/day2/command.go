package day2

import (
	"log"
	"strconv"
	"strings"
)

type CommandType int

type Command struct {
	Type  CommandType
	Value int
}

const (
	Forward CommandType = iota
	Down
	Up
)

func processCommands(commands []Command, position Position) int {
	for _, command := range commands {
		handleCommand(position, &command)
	}

	return position.calculateResult()
}

func handleCommand(position Position, command *Command) {
	switch command.Type {
	case Forward:
		position.moveForward(command.Value)
	case Down:
		position.moveDown(command.Value)
	case Up:
		position.moveUp(command.Value)
	}
}

func parseCommand(line string) Command {
	parts := strings.Split(line, " ")

	var commandType CommandType
	switch parts[0] {
	case "forward":
		commandType = Forward
	case "down":
		commandType = Down
	case "up":
		commandType = Up
	default:
		log.Fatalf("Couldn't parse %s as command", parts[0])
	}

	num, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Couldn't parse %s as number", parts[1])
	}

	return Command{
		Type:  commandType,
		Value: num,
	}
}

func parseCommands(input []string) (result []Command) {
	for _, line := range input {
		result = append(result, parseCommand(line))
	}
	return
}
