package main

import (
	"strconv"

	"github.com/brandonstinson/advent-of-code-2021/utils"
)

func main() {
	lines := utils.SplitFileByLine("d1/input.txt")

	utils.Print(1, partOne(lines), partTwo(lines))
}

func partOne(lines []string) int {
	var increases int

	for i, v := range lines {
		if i > 0 {
			current, err := strconv.Atoi(v)
			utils.ErrorCheck(err, "Error converting string to int")
			prev, err := strconv.Atoi(lines[i-1])
			utils.ErrorCheck(err, "Error converting string to int")

			if current > prev {
				increases++
			}
		}
	}
	return increases
}

func partTwo(lines []string) int {
	var increases int

	for i := 3; i < len(lines); i++ {
		one, err := strconv.Atoi(lines[i-3])
		utils.ErrorCheck(err, "Error converting string to int")
		two, err := strconv.Atoi(lines[i-2])
		utils.ErrorCheck(err, "Error converting string to int")
		three, err := strconv.Atoi(lines[i-1])
		utils.ErrorCheck(err, "Error converting string to int")
		four, err := strconv.Atoi(lines[i])
		utils.ErrorCheck(err, "Error converting string to int")

		if two+three+four > one+two+three {
			increases++
		}
	}

	return increases
}
