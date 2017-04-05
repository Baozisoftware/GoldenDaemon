package GoldenDaemon

import (
	"os"
	"errors"
	"os/exec"
)

var path = os.Args[0]
var args = os.Args[1:]

func parseArgs(key string) (b bool) {
	b = false
	for i, v := range args {
		if v == "-"+key || v == "-"+key+"=true" {
			b = true
			args = append(args[0:i], args[i+1:]...)
			break
		}
	}
	return
}

func RegisterTrigger(key string, otherArgs ...string) error {
	if parseArgs(key) {
		return RestartSelf(otherArgs...)
	}
	return errors.New("No Triggered.")
}

func RestartSelf(otherArgs ...string) error {
	if otherArgs != nil {
		args = append(args, otherArgs...)
	}
	cmd := exec.Command(path, args...)
	err := cmd.Start()
	if err != nil {
		return err
	}
	if cmd.Process != nil {
		os.Exit(0)
		return nil
	} else {
		return errors.New("Failed.")
	}
}
