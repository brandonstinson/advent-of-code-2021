package main

import (
	"strconv"

	"github.com/brandonstinson/advent-of-code-2021/utils"
)

func main() {
	lines := utils.SplitFileByLine("d01/input.txt")

	utils.Print(1, partOne(lines), partTwo(lines))
}

func partOne(lines []string) int {
	var increases int

	for i, v := range lines {
		if i > 0 {
			current, err := strconv.Atoi(v)
			utils.ErrorCheck(err)
			prev, err := strconv.Atoi(lines[i-1])
			utils.ErrorCheck(err)

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
		utils.ErrorCheck(err)
		two, err := strconv.Atoi(lines[i-2])
		utils.ErrorCheck(err)
		three, err := strconv.Atoi(lines[i-1])
		utils.ErrorCheck(err)
		four, err := strconv.Atoi(lines[i])
		utils.ErrorCheck(err)

		if two+three+four > one+two+three {
			increases++
		}
	}

	return increases
}
