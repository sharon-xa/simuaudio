package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ListAudioDevices() {
	outputStr, err := getDevices()
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(outputStr, "\n")
	devices := make(map[string]string)

	i := 0
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		nextLine := ""
		if i != len(lines)-1 {
			nextLine = strings.TrimSpace(lines[i+1])
		}

		if strings.HasPrefix(line, "node.name") && strings.HasPrefix(nextLine, "media.name") {
			devices[extractValue(nextLine)] = extractValue(line)
		} else if strings.HasPrefix(line, "node.name") && (strings.HasPrefix(nextLine, "node.name") || nextLine == "") {
			devices["Built-In Audio (speakers or audio jack)"] = extractValue(line)
		}
		i++
	}

	if len(devices) <= 0 {
		fmt.Println("No device found")
		return
	}

	arrow := "\033[1;31m->\033[0m" // Red and bold
	for name, nodeName := range devices {
		fmt.Printf("%s %s %s\n", name, arrow, nodeName)
	}
}

func getDevices() (string, error) {
	cmd := exec.Command("pactl", "list", "sinks")
	find := exec.Command("grep", "-E", "node.name|media.name")

	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		return "", err
	}

	// Connect cmd's stdout to the pipe's writer
	cmd.Stdout = pipeWriter

	// Connect find's stdin to the pipe's reader
	find.Stdin = pipeReader

	// Connect find's stdout to a byte buffer
	var output bytes.Buffer
	find.Stdout = &output

	// Start both commands
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := find.Start(); err != nil {
		panic(err)
	}

	// Close the writer when "cmd" finishes to signal EOF to "find"
	go func() {
		defer pipeWriter.Close()
		cmd.Wait()
	}()

	// Wait for "cmd" to finish
	find.Wait()

	outputStr := strings.TrimSpace(output.String())
	return outputStr, nil
}

// Helper function to extract value from a line like `key = "value"`
func extractValue(line string) string {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) < 2 {
		return ""
	}
	return strings.Trim(parts[1], " \"")
}
