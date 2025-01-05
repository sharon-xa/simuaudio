package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/sharon-xa/simuaudio/cmd"
)

func main() {
	start := flag.String(
		"s",
		"",
		"Play audio simultaneously between devices. Specify devices as <device1>=<latency1>,<device2>=<latency2>,...",
	)
	kill := flag.Bool("k", false, "Kill all simultaneous audio playbacks")
	list := flag.Bool("l", false, "List all audio devices")
	flag.Parse()

	switch {
	case *start != "":
		deviceLatencies := parseDeviceLatencies(*start)
		cmd.StartSimuAudio(deviceLatencies)

	case *kill:
		cmd.KillSimuAudio()

	case *list:
		cmd.ListAudioDevices()

	default:
		fmt.Println(
			"Usage: simuaudio -s <device1>=<latency1>,<device2>=<latency2> OR simuaudio -k OR simuaudio -l",
		)
	}
}

// parseDeviceLatencies parses the input string to extract device names and their latencies.
// If a latency is not provided, it defaults to 0ms.
func parseDeviceLatencies(input string) map[string]int {
	deviceLatencies := make(map[string]int)
	devices := strings.Split(input, ",")

	for _, device := range devices {
		parts := strings.Split(device, "=")
		deviceName := parts[0]
		latency := 0

		if len(parts) > 1 {
			parsedLatency, err := strconv.Atoi(parts[1])
			if err == nil {
				latency = parsedLatency
			} else {
				fmt.Printf("Warning: Invalid latency for device %s. Using default latency of 0ms.\n", deviceName)
			}
		}
		deviceLatencies[deviceName] = latency
	}

	return deviceLatencies
}
