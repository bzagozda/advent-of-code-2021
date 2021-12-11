package main

import (
	"log"

	"github.com/bzagozda/aoc2021/day"
	"github.com/bzagozda/aoc2021/day/day1"
	"github.com/bzagozda/aoc2021/day/day10"
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
	log.Println(day.Run(&day1.Day1{}))
	log.Println(day.Run(&day2.Day2{}))
	log.Println(day.Run(&day3.Day3{}))
	log.Println(day.Run(&day4.Day4{}))
	log.Println(day.Run(&day5.Day5{}))
	log.Println(day.Run(&day6.Day6{}))
	log.Println(day.Run(&day7.Day7{}))
	log.Println(day.Run(&day8.Day8{}))
	log.Println(day.Run(&day9.Day9{}))
	log.Println(day.Run(&day10.Day10{}))
}
