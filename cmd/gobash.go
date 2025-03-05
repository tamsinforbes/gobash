package cmd

import (
	"fmt"
	"os/exec"
)



func List(relative_path string) (result string) {
	// subpath := "./" + subfolder
	lst := *exec.Command("ls", relative_path)
	// The `Output` method executes the command and
	// collects the output, returning its value
	out, err := lst.Output()
	if err != nil {
  	// if there was any error, print it here
  	fmt.Println("could not run command: ", err)
		return
	}
	// otherwise, print the output from running the command
  return fmt.Sprint(string(out))
}