package utils

import (
	"os/exec"
)

func GetIP() string {
	curl := exec.Command("curl", "ip.sb")
	out, _ := curl.Output()
	return string(out)
}
