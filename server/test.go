package server

import (
	"errors"
	"fmt"
	"os/exec"
)

type TestArgs struct {
	Packages []string
}

type TestResults struct {
	Output []byte
	Err string
}

type Test int

func (t *Test) Runner(args *TestArgs, results *TestResults) error {
	if len(args.Packages) == 0 {
		return errors.New("no packages to install")
	}
	command := []string{"-al"}
	command = append(command, args.Packages...)
	out, err := exec.Command("ls", command...).CombinedOutput()
	if err != nil {
		results.Err = err.Error()
		fmt.Println("Runner Error [ " + results.Err + " ]")
	} else {
		fmt.Println("Successfully ran the runner...")
	}
	results.Output = out
	return nil
}
