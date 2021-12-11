package main

import (
	"log"

	"github.com/bzagozda/aoc2021/day"
	"github.com/bzagozda/aoc2021/day/day1"
	"github.com/bzagozda/aoc2021/day/day10"
	"github.com/bzagozda/aoc2021/day/day11"
	"github.com/bzagozda/aoc2021/day/day2"
	"github.com/bzagozda/aoc2021/day/day3"
	"github.com/bzagozda/aoc2021/day/day4"
	"github.com/bzagozda/aoc2021/day/day5"
	"github.com/bzagozda/aoc2021/day/day6"
	"github.com/bzagozda/aoc2021/day/day7"
	"github.com/bzagozda/aoc2021/day/day8"
	"github.com/bzagozda/aoc2021/day/day9"
)

func main() {
	log.Println(day.Run(&day1.Day{}))
	log.Println(day.Run(&day2.Day{}))
	log.Println(day.Run(&day3.Day{}))
	log.Println(day.Run(&day4.Day{}))
	log.Println(day.Run(&day5.Day{}))
	log.Println(day.Run(&day6.Day{}))
	log.Println(day.Run(&day7.Day{}))
	log.Println(day.Run(&day8.Day{}))
	log.Println(day.Run(&day9.Day{}))
	log.Println(day.Run(&day10.Day{}))
	log.Println(day.Run(&day11.Day{}))
}
