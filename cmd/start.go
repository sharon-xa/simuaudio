package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/sharon-xa/simuaudio/storage"
)

type Data struct {
	ModuleId  string            `json:"module_id"`
	DevicesId map[string]string `json:"devices_id"`
}

func StartSimuAudio(devices map[string]int) {
	// check if we have a connection already
	fs, err := storage.NewFileStorage("simuaudio", "data.json")
	if err != nil {
		fmt.Println(err)
		fmt.Println("couldn't save the module ids.")
		fmt.Println("you might want to unload modules manually")
		fmt.Println("\nyou can unload with the following command:")
		fmt.Println("pactl unload-module module_id")
		return
	}

	var data Data
	fs.Load(&data)
	if data.ModuleId != "" || (data.DevicesId != nil && len(data.DevicesId) > 0) {
		fmt.Println(
			"make sure to kill the connection you created before starting another connection",
		)
		return
	}

	cmd := exec.Command(
		"pactl",
		"load-module",
		"module-null-sink",
		"sink_name=combined_sink",
	)

	combinedSinkIDBytes, err := cmd.Output()
	if err != nil {
		fmt.Println("Error loading module-null-sink:", err)
		return
	}

	combinedSinkID := strings.TrimSpace(string(combinedSinkIDBytes))

	var idData Data
	idData.ModuleId = combinedSinkID
	idData.DevicesId = make(map[string]string)

	for device, latency := range devices {
		cmd := exec.Command(
			"pactl",
			"load-module",
			"module-loopback",
			"source=input.combined_sink.monitor",
			fmt.Sprintf("sink=\"%s\"", device),
			fmt.Sprintf("latency_msec=%d", latency),
		)

		deviceSinkIDBytes, err := cmd.Output()
		if err != nil {
			fmt.Println("Error loading module-loopback:", err)
			return
		}

		idData.DevicesId[device] = strings.TrimSpace(string(deviceSinkIDBytes))
	}

	cmd = exec.Command("pactl", "set-default-sink", "combined_sink")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error setting default sink to combined_sink:", err)
		return
	} else {
		fmt.Println("combined_sink set as default audio output")
	}

	err = fs.Save(idData)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Couldn't save the module ids.")
		fmt.Println("You might want to unload modules manually")
		fmt.Println("IDs: \n", idData.ModuleId)
		for _, id := range idData.DevicesId {
			fmt.Println(id)
		}
		fmt.Println("\nyou can unload with the following command:")
		fmt.Println("pactl unload-module module_id")
		return
	}

	fmt.Println("Simultaneous audio playback started")
}
