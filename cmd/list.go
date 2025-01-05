package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ListAudioDevices() {
	cmd := exec.Command("pactl", "list", "sinks")
	find := exec.Command("grep", "-E", "node.name|media.name")

	pipeReader, pipeWriter, err := os.Pipe()
	if err != nil {
		fmt.Println("PIPE Error:", err)
		return
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

	// Close the writer when cmd finishes to signal EOF to find
	go func() {
		defer pipeWriter.Close()
		cmd.Wait()
	}()

	// Wait for cmd to finish
	find.Wait()

	outputStr := strings.TrimSpace(output.String())
	lines := strings.Split(outputStr, "\n")

	devices := make(map[string]string)
	var currentNode string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "node.name") {
			currentNode = extractValue(line)
		} else if strings.HasPrefix(line, "media.name") {
			mediaName := extractValue(line)
			if currentNode != "" {
				devices[mediaName] = currentNode
				currentNode = "" // reset
			}
		}
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

// Helper function to extract value from a line like `key = "value"`
func extractValue(line string) string {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) < 2 {
		return ""
	}
	return strings.Trim(parts[1], " \"")
}
