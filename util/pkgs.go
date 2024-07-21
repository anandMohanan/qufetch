package util

import (
	"os/exec"
	"strconv"
	"strings"
)

// packages count installed on the current system
func Pkgs() string {
	cmds := map[string]string{
		"dpkg":   "dpkg-query -f '${binary:Package}\n' -W",
		"rpm":    "rpm -qa",
		"pacman": "pacman -Q",
	}

	for cmd, query := range cmds {
		if _, err := exec.LookPath(cmd); err == nil {
			out, err := exec.Command("sh", "-c", query).Output()
			if err != nil {
				return "error"
			}
			lines := strings.Split(string(out), "\n")
			count := len(lines) - 1
			return strconv.Itoa(count)
		}
	}

	return "0"
}
