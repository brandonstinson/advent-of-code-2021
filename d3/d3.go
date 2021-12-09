package main

import (
	"github.com/brandonstinson/advent-of-code-2021/utils"
)

func main() {
	lines := utils.SplitFileByLine("d3/input.txt")

	utils.Print(3, partOne(lines), partTwo(lines))
}

func partOne(lines []string) int {

	inverted := invert(lines)

	var gamma string
	var epsilon string

	for _, num := range inverted {
		value, _ := getMostAndLeastCommonBinaryDigit(num)
		if value == "1" {
			gamma += "1"
			epsilon += "0"
		}
		if value == "0" {
			gamma += "0"
			epsilon += "1"
		}
	}

	return utils.ConvertBinaryStringToInt(gamma) * utils.ConvertBinaryStringToInt(epsilon)
}

func partTwo(lines []string) int {

	o2, co2 := findLast(lines)

	return utils.ConvertBinaryStringToInt(o2) * utils.ConvertBinaryStringToInt(co2)
}

func invert(lines []string) []string {
	shifted := make([]string, len(lines[0]))

	for _, l := range lines {
		for i, v := range l {
			shifted[i] += string(v)
		}
	}

	return shifted
}

func getMostAndLeastCommonBinaryDigit(s string) (string, string) {
	var ones int
	var zeroes int
	for _, c := range s {
		if string(c) == "0" {
			zeroes++
		}
		if string(c) == "1" {
			ones++
		}
	}

	var most string
	var least string

	if ones > zeroes {
		most = "1"
		least = "0"
	} else if zeroes > ones {
		most = "0"
		least = "1"
	} else if zeroes == ones {
		most = "e"
		least = "e"
	}

	return most, least
}

func findLast(s []string) (string, string) {

	var inverted []string

	o2 := make([]string, len(s))
	copy(o2, s)

	o2Depth := 0

	for len(o2) > 1 {
		inverted = invert(o2)
		m, _ := getMostAndLeastCommonBinaryDigit(inverted[o2Depth])
		if m == "e" {
			o2 = filter(o2, o2Depth, "1")
		} else {
			o2 = filter(o2, o2Depth, m)
		}
		o2Depth++
	}

	co2 := make([]string, len(s))
	copy(co2, s)

	co2Depth := 0

	for len(co2) > 1 {
		inverted = invert(co2)
		_, l := getMostAndLeastCommonBinaryDigit(inverted[co2Depth])
		if l == "e" {
			co2 = filter(co2, co2Depth, "0")
		} else {
			co2 = filter(co2, co2Depth, l)
		}
		co2Depth++
	}

	return o2[0], co2[0]
}

func filter(s []string, i int, v string) []string {
	filtered := []string{}

	for _, x := range s {
		if string(x[i]) == v {
			filtered = append(filtered, x)
		}
	}

	return filtered
}
