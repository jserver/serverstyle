package server

import (
	"bytes"
	"os/exec"
)

type AptUpdateArgs struct{}

type AptUpdateResults struct {
	Err    string
	Output []byte
	Errors []byte
}

func (a AptUpdateResults) GetErr() string {
	return a.Err
}

func (a AptUpdateResults) GetOutput() string {
	return string(a.Output)
}

func (a AptUpdateResults) GetErrors() string {
	return string(a.Errors)
}

type AptUpdate struct{}

func (t *AptUpdate) Update(args *AptUpdateArgs, results *AptUpdateResults) error {
	cmd := exec.Command("apt-get", "-qq", "-y", "update")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		results.Err = err.Error()
		logger.Println("AptGet Update Error [ " + results.Err + " ]")
	} else {
		logger.Println("Successfully ran AptGet Update...")
	}
	results.Output = stdout.Bytes()
	results.Errors = stderr.Bytes()
	return nil
}
