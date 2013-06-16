package server

import (
	"os/exec"
)

type AptUpgradeArgs struct{}

type AptUpgradeResults struct {
	Err    string
	Output []byte
}

func (a AptUpgradeResults) GetErr() string {
	return a.Err
}

func (a AptUpgradeResults) GetOutput() string {
	return string(a.Output)
}

type AptUpgrade struct{}

func (t *AptUpgrade) Upgrade(args *AptUpgradeArgs, results *AptUpgradeResults) error {
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
