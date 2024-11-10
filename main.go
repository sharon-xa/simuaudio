package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	start := flag.Bool(
		"s",
		false,
		"Play audio simultaneously between devices\nput the name of the device that you want to play sound from without the need for the device that you're already playing sound from",
	)
	kill := flag.Bool("k", false, "Kill all simultaneous audio playbacks")
	list := flag.Bool("l", false, "List all audio devices")
	flag.Parse()

	switch {
	case *start:
		startSimuAudio()

	case *kill:
		killSimuAudio()

	case *list:
		listAudioDevices()

	default:
		fmt.Println(
			"Usage: simuaudio -s <device1> <device2> ... OR simuaudio -k ... OR simuaudio -l",
		)
	}
}

func startSimuAudio() {
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: simuaudio -s <device1> <device2> ...")
		os.Exit(1)
	}

	devices := flag.Args()
	cmd := exec.Command(
		"pactl",
		"load-module",
		"module-combine-sink",
		fmt.Sprintf("slaves=%s", strings.Join(devices, ",")),
	)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error loading module-combine-sink:", err)
		return
	}

	fmt.Printf(
		"Simultaneous audio playback started with module-combine-sink (Module ID: %s)\n",
		strings.TrimSpace(string(output)),
	)

	cmd = exec.Command("pactl", "set-default-sink", "combined")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error setting default sink to combined:", err)
	} else {
		fmt.Println("Combined output set as default.")
	}
}

func killSimuAudio() {
	cmd := exec.Command("pactl", "unload-module", "module-combine-sink")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Failed to stop simultaneous audio playback:", err)
	} else if strings.TrimSpace(string(output)) == "" {
		fmt.Println("Simultaneous audio playback stopped.")
	} else {
		fmt.Println("No module-combine-sink to kill")
	}
}

func listAudioDevices() {
	cmd := exec.Command("pactl", "list", "short", "sinks")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error listing audio devices:", err)
		return
	}

	fmt.Print(string(output))
}
