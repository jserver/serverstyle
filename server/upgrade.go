package server

import (
	"os/exec"
)

type AptGetUpgradeArgs struct {}

type AptGetUpgradeResults struct {
	Err    string
	Output []byte
}

func (a AptGetUpgradeResults) GetErr() string {
	return a.Err
}

func (a AptGetUpgradeResults) GetOutput() string {
	return string(a.Output)
}

type AptGetUpgrade struct {}

func (t *AptGetUpgrade) Upgrade(args *AptGetUpgradeArgs, results *AptGetUpgradeResults) error {
	command := []string{"apt-get", "-y", "upgrade"}
	out, err := exec.Command("sudo", command...).CombinedOutput()
	if err != nil {
		results.Err = err.Error()
		logger.Println("AptGet Upgrade Error [ " + results.Err + " ]")
	} else {
		logger.Println("Successfully ran AptGet Upgrade...")
	}
	results.Output = out
	return nil
}
