package day2

type Day2 struct{}

func (d *Day2) Number() int {
	return 2
}

func (d *Day2) Part1(input []string) int {
	return processCommands(parseCommands(input), &PositionBasic{})
}

func (d *Day2) Part2(input []string) int {
	return processCommands(parseCommands(input), &PositionAdvanced{})
}
