package server

import (
	"bytes"
	"errors"
	"os/exec"
)

type TestArgs struct {
	Auth
	Dirs []string
}

type TestResults struct {
	Err    string
	Output []byte
	Errors []byte
}

func (t TestResults) GetErr() string {
	return t.Err
}

func (t TestResults) GetOutput() string {
	return string(t.Output)
}

func (t TestResults) GetErrors() string {
	return string(t.Errors)
}

type Test struct{}

func (t *Test) Runner(args *TestArgs, results *TestResults) error {
	if args.Username != "user" || args.Password != "pass123" {
		err := errors.New("Auth Failed")
		logger.Println(err.Error())
		return err
	}

	if len(args.Dirs) == 0 {
		err := errors.New("No directories to display")
		logger.Println(err.Error())
		return err
	}
	command := []string{"-al"}
	command = append(command, args.Dirs...)
	cmd := exec.Command("ls", command...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		results.Err = err.Error()
		logger.Println("Runner Error [ " + results.Err + " ]")
	} else {
		logger.Println("Successfully ran the runner...")
	}
	results.Output = stdout.Bytes()
	results.Errors = stderr.Bytes()
	return nil
}
