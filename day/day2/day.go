package day2

type Day struct{}

func (d *Day) Number() int {
	return 2
}

func (d *Day) Part1(input []string) int {
	return processCommands(parseCommands(input), &PositionBasic{})
}

func (d *Day) Part2(input []string) int {
	return processCommands(parseCommands(input), &PositionAdvanced{})
}
