package server

import (
	"bytes"
	"errors"
	"os/exec"
)

type AptInstallArgs struct {
	Packages []string
}

type AptInstallResults struct {
	Err    string
	Output []byte
	Errors []byte
}

func (a AptInstallResults) GetErr() string {
	return a.Err
}

func (a AptInstallResults) GetOutput() string {
	return string(a.Output)
}

func (a AptInstallResults) GetErrors() string {
	return string(a.Errors)
}

type AptInstall struct{}

func (t *AptInstall) Install(args *AptInstallArgs, results *AptInstallResults) error {
	if len(args.Packages) == 0 {
		return errors.New("no packages to install")
	}
	command := []string{"DEBIAN_FRONTEND=noninteractive", "apt-get", "-y", "install"}
	command = append(command, args.Packages...)
	cmd := exec.Command("env", command...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		results.Err = err.Error()
		logger.Println("AptGet Install Error [ " + results.Err + " ]")
	} else {
		logger.Println("Successfully ran AptGet Install...")
	}
	results.Output = stdout.Bytes()
	results.Errors = stderr.Bytes()
	return nil
}
