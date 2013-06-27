package server

import (
	"bytes"
	"errors"
	"os/exec"
)

type EasyInstallArgs struct {
	Packages []string
}

type EasyInstallResults struct {
	Err    string
	Output []byte
	Errors []byte
}

func (e EasyInstallResults) GetErr() string {
	return e.Err
}

func (e EasyInstallResults) GetOutput() string {
	return string(e.Output)
}

func (e EasyInstallResults) GetErrors() string {
	return string(e.Errors)
}

type EasyInstall struct{}

func (t *EasyInstall) Install(args *EasyInstallArgs, results *EasyInstallResults) error {
	if len(args.Packages) == 0 {
		return errors.New("no packages to install")
	}
	command := []string{"install"}
	command = append(command, args.Packages...)
	cmd := exec.Command("pip", command...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		results.Err = err.Error()
		logger.Println("Easy Install Error [ " + results.Err + " ]")
	} else {
		logger.Println("Successfully ran Easy Install...")
	}
	results.Output = stdout.Bytes()
	results.Errors = stderr.Bytes()
	return nil
}
