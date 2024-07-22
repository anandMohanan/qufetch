package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Memory() string {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return "error"
	}

	lines := strings.Split(string(data), "\n")

	var memTotal, memFree, memAvailable uint64
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "MemTotal:":
			memTotal, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return "error"
			}
		case "MemFree:":
			memFree, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return "error"
			}
		case "MemAvailable:":
			memAvailable, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return "error"
			}
		}
	}

	usedKB := memTotal - memAvailable
	if memAvailable == 0 {
		usedKB = memTotal - memFree
	}

	totalMB := memTotal / 1024
	usedMB := usedKB / 1024

	return fmt.Sprintf("%dM / %dM", usedMB, totalMB)
}
