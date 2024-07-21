package util

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// OS returns the name of the operating system
func OS() string {
	switch runtime.GOOS {
	case "linux":
		return getLinuxOS()
	case "darwin":
		return getMacOS()
	default:
		return "Unsupported OS"
	}
}

// getLinuxOS retrieves the Linux distribution name from /etc/os-release
func getLinuxOS() string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "Error: " + err.Error()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			osName := strings.TrimPrefix(line, "PRETTY_NAME=")
			return strings.Trim(osName, `"`) // Remove surrounding quotes
		}
	}

	if err := scanner.Err(); err != nil {
		return "Error: " + err.Error()
	}

	return "Unknown Linux distribution"
}

// getMacOS retrieves the macOS name using the sw_vers command
func getMacOS() string {
	out, err := exec.Command("sw_vers", "-productName").Output()
	if err != nil {
		return "Error: " + err.Error()
	}
	name := strings.TrimSpace(string(out))

	versionOut, err := exec.Command("sw_vers", "-productVersion").Output()
	if err != nil {
		return "Error: " + err.Error()
	}
	version := strings.TrimSpace(string(versionOut))

	return name + " " + version
}
