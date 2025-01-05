# SimuAudio

SimuAudio is a Command-Line Interface (CLI) tool for managing simultaneous audio playback between multiple devices using PipeWire's module-combine-sink. It allows you to play audio through multiple devices, stop simultaneous playback, and list all available audio devices.

## Features

- Start simultaneous audio playback between devices.
- Combine multiple audio output devices into a single virtual output.
- Kill all active playback processes.
- List all available audio devices.
- Automatically set the combined output as the default for playback.
- Check if a device is already active in a loopback before starting a new one.

## Installation

You can manually install the `SimuAudio` CLI by cloning this repository and building it yourself, or you can use an installation script.

### Installation via Script

Run the following command in your terminal to install SimuAudio:

```bash
curl -fsSL https://raw.githubusercontent.com/sharon-xa/simuaudio/main/install.sh | sudo bash
```

## Dependencies
1. pipewire
2. pipewire-pulse (called pipewire-pulseaudio on fedora)
3. pipewire-alsa
4. wireplumber
5. pipewire-utils pulseaudio-utils


## Usage

### Command Line Usage

```bash
simuaudio [flags] [device...]
```

### Flags

| Flag | Description                                                                  |
| ---- | ---------------------------------------------------------------------------- |
| `-s` | Start simultaneous audio playback between devices.                           |
| `-k` | Kill all active playback processes, stopping audio playback between devices. |
| `-l` | List all available audio devices.                                            |
| `-h` | Show the help message with available operations.                             |

### Example Usage

1. **Starting Simultaneous Audio Playback**
   To start audio playback across multiple devices:
   
   ```bash
   simuaudio -s <device1> <device2>
   ```
   
   For example:
   
   ```bash
   simuaudio -s alsa_output.pci-0000_0a_00.6.analog-stereo bluez_output.D8_37_3B_09_3D_C0.1
   ```
   
   This will create a combined output that plays the same audio on both devices. The CLI will also set the combined output as the default, so all system audio is routed to both devices simultaneously.

2. **Killing Active Loopbacks**
   To stop all active playback processes, run:
   
   ```bash
   simuaudio -k
   ```
   
   This will unload the `module-combine-sink` module, effectively stopping all simultaneous audio playback.

3. **Listing Audio Devices**
   To list all available audio devices (sinks):
   
   ```bash
   simuaudio -l
   ```
   
   This command will display all available audio output devices on your system, helping you identify the correct device names to use with `-s`.

### Additional Notes

- The CLI automatically uses `module-combine-sink` to combine the specified devices, avoiding the need for manual device activity checks.
- Devices must be specified correctly for playback to work. Make sure the device names (e.g., `alsa_output.pci-0000_0a_00.6.analog-stereo`) are valid and available in your system.
- If a combined output is created, it will be set as the default sink for your audio, allowing all applications to play through both devices without additional configuration.
  
### Troubleshooting

If you encounter issues with **SimuAudio**, here are some steps that might help:

1. **Ensure PipeWire and Dependencies are Installed**: Verify that PipeWire, `pactl`, and `pw-jack` are properly installed and running on your system. You can check the status of PipeWire with:
   ```bash
   systemctl --user status pipewire
   ```

2. **Volume and Device Configuration**: If some devices aren’t producing sound, check your audio configuration with `pavucontrol` (PulseAudio Volume Control) to ensure that the correct devices are selected as output for the combined sink. 

3. **Using QjackCtl**: For more advanced troubleshooting and visualization, you can use **QjackCtl**, a graphical tool that displays audio routing and allows manual control over connections. This can be helpful if you need to see how audio flows between devices and to troubleshoot routing issues.

   - Install **QjackCtl**:
     ```bash
     sudo apt install qjackctl
     ```
   - Launch it and view the connections between devices. This can reveal if any devices are improperly routed or disconnected.

4. **Reboot PipeWire**: Sometimes, restarting PipeWire can resolve audio issues:
   ```bash
   systemctl --user restart pipewire
   ```

5. **Kill All Combined Sinks**: If multiple combined sinks are causing conflicts, you can remove them all with:
   ```bash
   pactl unload-module module-combine-sink
   ```

6. **Device Name Mismatch**: Double-check the device names listed with `simuaudio -l` to ensure you’re using correct names when specifying devices.

  
This CLI provides a convenient way to manage audio output on multiple devices, leveraging PipeWire’s `module-combine-sink` for seamless audio sharing.

