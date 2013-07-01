package server

import (
	"bytes"
	"os/exec"
)

type PPAInstallArgs struct {
	Name, Package string
}

type PPAInstallResults struct {
	Err    string
	Output []byte
	Errors []byte
}

func (a PPAInstallResults) GetErr() string {
	return a.Err
}

func (a PPAInstallResults) GetOutput() string {
	return string(a.Output)
}

func (a PPAInstallResults) GetErrors() string {
	return string(a.Errors)
}

type PPAInstall struct{}

func (t *PPAInstall) AddRepo(args *PPAInstallArgs, results *PPAInstallResults) error {
	cmd := exec.Command("add-apt-repository", "-y", "ppa:" + args.Name)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		results.Err = err.Error()
		logger.Println("PPA Install Error [ " + results.Err + " ]")
	} else {
		logger.Println("Successfully ran PPA Install...")
	}
	results.Output = stdout.Bytes()
	results.Errors = stderr.Bytes()
	return nil
}
