package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Uptime retrieves the system uptime in a human-readable format
func Uptime() string {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "Error: " + err.Error()
	}

	fields := strings.Fields(string(data))
	if len(fields) < 1 {
		return "Error: unexpected format of /proc/uptime"
	}

	uptimeSeconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return "Error: " + err.Error()
	}

	uptimeDuration := time.Duration(uptimeSeconds) * time.Second
	return fmtDuration(uptimeDuration)
}

// fmtDuration formats a time.Duration into a human-readable string
func fmtDuration(d time.Duration) string {
	days := d / (24 * time.Hour)
	d -= days * 24 * time.Hour
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	d -= minutes * time.Minute
	seconds := d / time.Second

	return fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds", days, hours, minutes, seconds)
}
