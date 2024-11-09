package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func checkSubstrings(str string, subs ...string) bool {
	isCompleteMatch := true

	for _, sub := range subs {
		if !strings.Contains(str, sub) {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch
}

// Function to check if loopback is already active for a given device
func isDeviceActive(device string) bool {
	cmd := exec.Command("pactl", "list", "short", "sinks")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error checking loopback status:", err)
		return false
	}
	o := string(output)
	outputList := strings.Split(o, "\n")

	for _, line := range outputList {
		if checkSubstrings(line, device, "RUNNING") {
			return true
		}
	}

	return false
}
