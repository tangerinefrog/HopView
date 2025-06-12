package commands

import "os/exec"

func Execute(command string, arg ...string) (string, error) {
	cmd := exec.Command(command, arg...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
