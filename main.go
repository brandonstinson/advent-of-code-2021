package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/brandonstinson/advent-of-code-2021/utils"
)

func main() {

	if len(os.Args) < 2 {
		utils.ErrorText("You must provide the day")
		return
	}

	day := os.Args[1]

	new := ""

	if len(os.Args) == 3 {
		new = os.Args[2]
	}

	if new == "new" {
		utils.NewDay(day)
	} else {
		_, err := os.Stat(fmt.Sprintf("d%s/d%s.go", day, day))
		utils.ErrorCheck(err, "No files exist for that day")
		out, err := exec.Command("go", "run", fmt.Sprintf("d%s/d%s.go", day, day)).Output()
		utils.ErrorCheck(err, "Error executing command")

		fmt.Printf("%s", string(out))
	}
}
