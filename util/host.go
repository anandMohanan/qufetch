package util

import (
	"os"
)

// host name
func Host() string {
	host := os.Getenv("HOSTNAME")
	if host == "" {
		host = os.Getenv("HOST")
	}
	if host == "" {
		host = os.Getenv("COMPUTERNAME")
	}
	if host == "" {
		host = os.Getenv("USERDOMAIN")
	}
	if host == "" {
		hostname, err := os.Hostname()
		if err == nil {
			host = hostname
		}
	}
	// if host is still empty, use uname
	if host == "" {
		host = os.Getenv("MACHINE")
	}
	// if host is still empty, use os.Getenv("USER")
	if host == "" {
		host = os.Getenv("USER")
	}
	// if host is still empty, use os.Getenv("LOGNAME")
	if host == "" {
		host = os.Getenv("LOGNAME")
	}
	// if host is still empty, use os.Getenv("USERNAME")
	if host == "" {
		host = os.Getenv("USERNAME")
	}
	return host
}
