package server

import (
	"errors"
	"os/exec"
)

type TestArgs struct {
	Dirs []string
}

type TestResults struct {
	Err    string
	Output []byte
}

func (t TestResults) GetErr() string {
	return t.Err
}

func (t TestResults) GetOutput() string {
	return string(t.Output)
}

type Test struct{}

func (t *Test) Runner(args *TestArgs, results *TestResults) error {
	if len(args.Dirs) == 0 {
		return errors.New("No directories to display")
	}
	command := []string{"-al"}
	command = append(command, args.Dirs...)
	out, err := exec.Command("ls", command...).CombinedOutput()
	if err != nil {
		results.Err = err.Error()
		logger.Println("Runner Error [ " + results.Err + " ]")
	} else {
		logger.Println("Successfully ran the runner...")
	}
	results.Output = out
	return nil
}
