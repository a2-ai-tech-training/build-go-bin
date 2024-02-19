package main

import (
	"fmt"
	"log"
	//"os"
	"os/exec"
)

func main() {
	//cmd := "/bin/ls /"
	args := ["/bin/ls", "~"]
	cmd2 := exec.Command(args)
	_, err := cmd2.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}

	//fmt.Println(os.Environ())
	fmt.Printf("Standard Input:\n%s\n", cmd2.Args)
	//fmt.Println(cmd2.Stdin)
	fmt.Printf("Standard Output:\n%s\n", cmd2.Stdout)
	fmt.Printf("Standard Error:\n%s\n", cmd2.Stderr)
	fmt.Printf("Exit Code:\n%d\n", cmd2.ProcessState.ExitCode())
	//fmt.Println(cmd2.ProcessState.ExitCode())
}
