package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/sharon-xa/simuaudio/storage"
)

func KillSimuAudio() {
	fs, err := storage.NewFileStorage("simuaudio", "data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var idData Data
	err = fs.Load(&idData)
	if err != nil {
		fmt.Println(err)
		return
	}

	if idData.ModuleId != "" {
		cmd := exec.Command("pactl", "unload-module", idData.ModuleId)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Failed to stop simultaneous audio playback:", err)
		} else if strings.TrimSpace(string(output)) == "" {
			fmt.Println("module-null-sink is unloaded")
		}
	} else {
		fmt.Println("No module-null-sink to kill")
	}

	if len(idData.DevicesId) <= 0 {
		fmt.Println("No devices found.")
		return
	}

	for device, id := range idData.DevicesId {
		cmd := exec.Command("pactl", "unload-module", id)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf(
				"Failed to stop simultaneous audio playback for (%s): %s\n",
				device,
				err.Error(),
			)
		} else if strings.TrimSpace(string(output)) == "" {
			fmt.Printf("Simultaneous audio playback stopped for (%s).\n", device)
		}
	}

	err = fs.Clear()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nsimultaneous audio stopped successfully")
}
