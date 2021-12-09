package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ErrorText(m string) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	fmt.Printf("%s%s%s\n", colorRed, m, colorReset)
}

func ErrorCheck(e error, m string) {
	if e != nil {
		ErrorText(m)
		log.Fatal(e)
	}
}

func SplitFileByLine(file string) []string {
	data, err := ioutil.ReadFile(file)
	ErrorCheck(err, "Error reading file")

	return strings.Split(string(data), "\n")
}

func Print(day int, one interface{}, two interface{}) {
	fmt.Printf("Day %d\nPart 1: %v\nPart 2: %v\n", day, one, two)
}

func ConvertBinaryStringToInt(s string) int {
	i, err := strconv.ParseInt(s, 2, 0)
	ErrorCheck(err, "Error converting string to int")
	return int(i)
}
