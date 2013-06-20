package server

import (
	"bytes"
	"os/exec"
)

type AptUpgradeArgs struct{}

type AptUpgradeResults struct {
	Err    string
	Output []byte
	Errors []byte
}

func (a AptUpgradeResults) GetErr() string {
	return a.Err
}

func (a AptUpgradeResults) GetOutput() string {
	return string(a.Output)
}

func (a AptUpgradeResults) GetErrors() string {
	return string(a.Errors)
}

type AptUpgrade struct{}

func (t *AptUpgrade) Upgrade(args *AptUpgradeArgs, results *AptUpgradeResults) error {
	command := []string{"apt-get", "-qq", "-y", "upgrade"}
	cmd := exec.Command("sudo", command...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		results.Err = err.Error()
		logger.Println("AptGet Upgrade Error [ " + results.Err + " ]")
	} else {
		logger.Println("Successfully ran AptGet Upgrade...")
	}
	results.Output = stdout.Bytes()
	results.Errors = stderr.Bytes()
	return nil
}
