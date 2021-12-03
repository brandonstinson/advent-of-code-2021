package main

import (
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/brandonstinson/advent-of-code-2021/utils"
)

func main() {
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		utils.ErrorCheck(err)
		if info.IsDir() && strings.HasPrefix(info.Name(), "d") {
			out, err := exec.Command("go", "run", fmt.Sprintf("%s/%s.go", path, path)).Output()
			utils.ErrorCheck(err)
			fmt.Printf("%s\n", string(out))
		}
		return nil
	})
	utils.ErrorCheck(err)
}
