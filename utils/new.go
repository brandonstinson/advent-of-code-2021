package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func NewDay(day string) {
	filePath := filepath.Join(fmt.Sprintf("d%s", day), fmt.Sprintf("d%s.go", day))
	if _, err := os.Stat(filePath); err == nil {
		ErrorText(fmt.Sprintf("Day %s already exists", day))
		os.Exit(1)
	} else {
		err := os.MkdirAll(filepath.Dir(filePath), 0777)
		ErrorCheck(err, "Error creating directory")
		f, err := os.Create(filePath)
		ErrorCheck(err, "Error creating file")

		defer f.Close()

		lines := []string{
			"package main\n\n",
			"func main() {\n",
			fmt.Sprintf("lines := utils.SplitFileByLine(\"d%s/input.txt\")\n\n", day),
			fmt.Sprintf("utils.Print(%s, partOne(lines), partTwo(lines))\n}\n\n", day),
			"func partOne(lines []string) int {\n",
			"return 1\n}\n\n",
			"func partTwo(lines []string) int {\n",
			"return 2\n}\n\n",
		}

		for _, v := range lines {
			_, err := f.WriteString(v)
			ErrorCheck(err, "Error writing string to file")
		}
	}
}
