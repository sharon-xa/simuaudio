package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	start := flag.Bool("s", false, "play audio simultaneously between devices\nput the name of the device that you want to play sound from without the need for the device that you're already playing sound from")
	kill := flag.Bool("k", false, "kill all simultaneous audio playbacks")
	list := flag.Bool("l", false, "list all audio devices")
	flag.Parse()

	switch {
	case *start:
		startSimuAudio()

	case *kill:
		killSimuAudio()

	case *list:
		listAudioDevices()

	default:
		fmt.Println("Usage: simuaudio -s <device1> <device2> ... OR simuaudio -k ... OR simuaudio -l")
	}
}

func startSimuAudio() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: simuaudio -s <device1> <device2> ...")
		os.Exit(1)
	}

	for _, device := range os.Args[2:] {
		if isDeviceActive(device) {
			fmt.Printf("Loopback already active for device %s\n", device)
			continue
		}

		cmd := exec.Command("pw-loopback",
			"--capture-props=node.name=loopback_capture",
			fmt.Sprintf("--playback-props=node.target=%s", device),
		)

		cmd.Stdout = nil
		cmd.Stderr = nil

		if err := cmd.Start(); err != nil {
			fmt.Printf("Error starting pw-loopback for device %s: %v\n", device, err)
		} else {
			fmt.Printf("Loopback started for device %s\n", device)
		}
	}
}

func killSimuAudio() {
	cmd := exec.Command("pkill", "pw-loopback")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to kill pw-loopback processes:", err)
	} else {
		fmt.Println("All pw-loopback processes killed successfully.")
	}
}

func listAudioDevices() {
	cmd := exec.Command("pactl", "list", "short", "sinks")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(output))
}
