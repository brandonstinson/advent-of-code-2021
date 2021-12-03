package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func ErrorCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func SplitFileByLine(file string) []string {
	data, err := ioutil.ReadFile(file)
	ErrorCheck(err)

	return strings.Split(string(data), "\n")
}

func Print(day int, one interface{}, two interface{}) {
	colorReset := "\033[0m"
	colorOne := "\033[33m" + "\033[40m"
	colorTwo := "\033[36m" + "\033[40m"
	colorThree := "\033[37;1m"
	d := fmt.Sprintf("%s%sDay %d%s%s", string(colorOne), strings.Repeat("-", 15), day, strings.Repeat("-", 15), string(colorReset))
	p1 := fmt.Sprintf("%sPart 1:%s %s%v%s", string(colorTwo), string(colorReset), string(colorThree), one, string(colorReset))
	p2 := fmt.Sprintf("%sPart 2:%s %s%v%s", string(colorTwo), string(colorReset), string(colorThree), two, string(colorReset))
	fmt.Printf("%s\n%s\n%s", d, p1, p2)
}
