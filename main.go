package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Solve all days
func main() {
	matches, _ := filepath.Glob("2024/day*/main.go")
	for _, fp := range matches {
		fmt.Printf("-- %v\n", fp)
		filename := filepath.Base(fp)
		cmd := exec.Command("go", "run", filename, "input.txt")
		cmd.Dir = filepath.Dir(fp)
		stdout, stderr := cmd.CombinedOutput()
		if len(stdout) > 0 {
			fmt.Printf("%v\n", string(stdout))
		}
		if stderr != nil {
			fmt.Fprintf(os.Stderr, "%v\n", stderr.Error())
		}
	}
}
