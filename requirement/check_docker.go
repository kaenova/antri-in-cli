package requirement

import (
	"os/exec"
)

func CheckDocker() (interface{}, string, string) {
	var pesanSukses string = ""
	var pesanError string = ""

	_, err := exec.Command("docker", "-v").Output()
	if err != nil {
		pesanError = err.Error()
		return nil, "", pesanError
	}
	pesanSukses = "Docker Terinstall"
	return nil, pesanSukses, ""
}
