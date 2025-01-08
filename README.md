# SimuAudio

SimuAudio is a Command-Line Interface (CLI) tool for managing simultaneous audio playback between multiple devices using PipeWire's `module-combine-sink`. It allows you to play audio through multiple devices, stop simultaneous playback, and list all available audio devices.

## Features

- Start simultaneous audio playback between devices, with optional latency settings.
- Combine multiple audio output devices into a single virtual output.
- Kill all active playback processes.
- List all available audio devices.
- Automatically set the combined output as the default for playback.
- Save and manage device configurations for ease of use.
- Validate device names and latency values.

## Installation

You can manually install the `SimuAudio` CLI by cloning this repository and building it yourself, or you can use an installation script.

### Installation via Script

Run the following command in your terminal to install SimuAudio:

```bash
curl -fsSL https://raw.githubusercontent.com/sharon-xa/simuaudio/main/install.sh | sudo bash
```

## Dependencies

1. PipeWire
2. pipewire-pulse (called pipewire-pulseaudio on Fedora)
3. pipewire-alsa
4. wireplumber
5. pipewire-utils
6. pulseaudio-utils

## Usage

### Command Line Usage

```bash
simuaudio [flags]
```

### Flags

| Flag | Description                                                                  |
| ---- | ---------------------------------------------------------------------------- |
| `-s` | Start simultaneous audio playback between devices with optional latencies.   |
| `-k` | Kill all active playback processes, stopping audio playback between devices. |
| `-l` | List all available audio devices.                                            |
| `-h` | Show the help message with available operations.                             |

### Example Usage

1. **Starting Simultaneous Audio Playback**

   To start audio playback across multiple devices with optional latencies:

   ```bash
   simuaudio -s <device1>=<latency1>,<device2>=<latency2>
   ```

   For example:

   ```bash
   simuaudio -s alsa_output.pci-0000_0a_00.6.analog-stereo=100,bluez_output.D8_37_3B_09_3D_C0.1=200
   ```

   This will create a combined output that plays the same audio on both devices, applying a latency of 100ms to the first device and 200ms to the second. If a latency value is omitted, it defaults to 0ms.

   The CLI will also set the combined output as the default, so all system audio is routed to the combined sink.

2. **Killing Active Loopbacks**

   To stop all active playback processes, run:

   ```bash
   simuaudio -k
   ```

   This will unload the `module-null-sink` and `module-loopback` modules, stopping all simultaneous audio playback.

3. **Listing Audio Devices**

   To list all available audio devices (sinks):

   ```bash
   simuaudio -l
   ```

   This command will display all available audio output devices on your system, helping you identify the correct device names and associated media names to use with the `-s` flag.

### Additional Notes

- Ensure device names are specified correctly for playback to work. Use the `-l` flag to view available devices.
- If a combined output is created, it will be set as the default sink, allowing all applications to play through the combined devices without additional configuration.
- If an invalid latency value is provided, a warning will be displayed, and a default latency of 0ms will be used for the affected device.

### Troubleshooting

If you encounter issues with **SimuAudio**, here are some steps that might help:

1. **Ensure PipeWire and Dependencies are Installed**: Verify that PipeWire, `pactl`, and `pw-jack` are properly installed and running on your system. You can check the status of PipeWire with:

   ```bash
   systemctl --user status pipewire
   ```

2. **Volume and Device Configuration**: If some devices aren’t producing sound, check your audio configuration with `pavucontrol` (PulseAudio Volume Control) to ensure that the correct devices are selected as outputs for the combined sink.

3. **Restart PipeWire**: Sometimes, restarting PipeWire can resolve audio issues:

   ```bash
   systemctl --user restart pipewire
   ```

4. **Unload All Modules**: If multiple combined sinks are causing conflicts, you can remove them all with:

   ```bash
   pactl unload-module module-combine-sink
   ```

5. **Verify Device Names**: Use `simuaudio -l` to list all available audio devices and ensure the device names used in `-s` are valid.

6. **Check Error Messages**: If a command fails, SimuAudio provides detailed error messages to help diagnose the issue.

This CLI simplifies the management of audio output on multiple devices, leveraging PipeWire’s advanced features for seamless audio sharing.


