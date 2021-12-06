package main

import (
	"strconv"
	"strings"

	"github.com/brandonstinson/advent-of-code-2021/utils"
)

func main() {
	lines := utils.SplitFileByLine("d2/input.txt")

	utils.Print(2, partOne(lines), partTwo(lines))
}

func partOne(lines []string) int {

	h := 0
	d := 0

	for _, l := range lines {
		parts := strings.Split(l, " ")
		dir := parts[0]
		amt, err := strconv.Atoi(parts[1])
		utils.ErrorCheck(err, "Error converting string to int")

		if dir == "forward" {
			h += amt
		}
		if dir == "down" {
			d += amt
		}
		if dir == "up" {
			d -= amt
		}
	}

	return d * h
}

func partTwo(lines []string) int {

	h := 0
	d := 0
	a := 0

	for _, l := range lines {
		parts := strings.Split(l, " ")
		dir := parts[0]
		amt, err := strconv.Atoi(parts[1])
		utils.ErrorCheck(err, "Error converting string to int")

		if dir == "forward" {
			h += amt
			d += a * amt
		}
		if dir == "down" {
			a += amt
		}
		if dir == "up" {
			a -= amt
		}
	}
	return d * h
}
