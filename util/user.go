package util

import (
	"fmt"
	"os"
	"os/user"
)

func User() string {
	u, err := user.Current()
	if err != nil {
		return "error"
	}

	hostname, err := os.Hostname()
	if err != nil {
		return "error"
	}

	return fmt.Sprintf("%s@%s", u.Username, hostname)
}
