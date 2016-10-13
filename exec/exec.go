package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "333", ">>", "abc")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error occurs")
	}
	print(string(stdout))
}
