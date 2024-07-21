package util

import (
	"strings"

	"golang.org/x/sys/unix"
)

// Kernel retrieves the kernel name and version
func Kernel() string {
	var uts unix.Utsname
	if err := unix.Uname(&uts); err != nil {
		return "Error: " + err.Error()
	}

	kernelName := trimNullBytes(uts.Sysname[:])
	kernelVersion := trimNullBytes(uts.Release[:])

	return kernelName + " " + kernelVersion
}

// trimNullBytes converts a null-terminated byte array to a string
func trimNullBytes(b []byte) string {
	return strings.TrimRight(string(b), "\x00")
}
