package day2

type Position interface {
	moveForward(int)
	moveUp(int)
	moveDown(int)
	calculateResult() int
}

type PositionBasic struct {
	horizontal int
	depth      int
}

func (p *PositionBasic) moveForward(distance int) {
	p.horizontal += distance
}

func (p *PositionBasic) moveUp(distance int) {
	p.depth -= distance
}

func (p *PositionBasic) moveDown(distance int) {
	p.depth += distance
}

func (p *PositionBasic) calculateResult() int {
	return p.depth * p.horizontal
}

type PositionAdvanced struct {
	horizontal int
	depth      int
	aim        int
}

func (p *PositionAdvanced) moveForward(distance int) {
	p.horizontal += distance
	p.depth += p.aim * distance
}

func (p *PositionAdvanced) moveUp(distance int) {
	p.aim -= distance
}

func (p *PositionAdvanced) moveDown(distance int) {
	p.aim += distance
}

func (p *PositionAdvanced) calculateResult() int {
	return p.depth * p.horizontal
}
