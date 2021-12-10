package day

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Day interface {
	Number() int
	Part1([]string) int
	Part2([]string) int
}

func readInput(day Day) []string {
	sessionToken := "53616c7465645f5f951341f371bc1613fd44080b25b0501973415f01c5be568693006035644ee6557555515f48601b64"
	filename := fmt.Sprintf("./input/%d.txt", day.Number())

	if _, err := os.Stat(filename); err != nil {
		est, err := time.LoadLocation("EST")
		if err != nil {
			panic(err)
		}

		if t := time.Date(2021, time.December, day.Number(), 0, 0, 0, 0, est); time.Until(t) > 0 {
			log.Fatal(fmt.Sprintf("Puzzle not unlocked yet! Sleeping for %d...\n", time.Until(t)))
		}

		log.Println("Downloading input...")
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%v/input", 2021, day.Number()), nil)
		req.AddCookie(&http.Cookie{
			Name:  "session",
			Value: sessionToken,
		})

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile(filename, data, 0660); err != nil {
			panic(err)
		}
	}

	return readFile(filename)
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func Run(day Day) (int, int) {
	input := readInput(day)
	return day.Part1(input), day.Part2(input)
}
