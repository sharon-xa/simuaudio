# SimuAudio

**SimuAudio** is a Command-Line Interface (CLI) tool for managing simultaneous audio playback between multiple devices using PipeWire's `pw-loopback` utility. It allows you to start audio playback between devices, kill active playback processes, and list all available audio devices.

## Features

- Start simultaneous audio playback between devices.
- Kill all active `pw-loopback` processes.
- List all available audio devices.
- Check if a device is already active in a loopback before starting a new one.

## Installation

You can manually install the `SimuAudio` CLI by cloning this repository and building it yourself, or you can use an installation script.

### Installation via Script (Linux only)

Run the following command in your terminal to install SimuAudio:

```bash
curl -fsSL https://raw.githubusercontent.com/sharon-xa/simuaudio/main/install.sh | sudo bash
```

## Usage

### Command Line Usage

```bash
simuaudio [flags] [device...]
```

### Flags

| Flag        | Description                                                         |
|-------------|---------------------------------------------------------------------|
| `-s`        | Start simultaneous audio playback between devices.                 |
| `-k`        | Kill all active `pw-loopback` processes.                           |
| `-l`        | List all available audio devices.                                  |
| `-h`        | Show the help message with available operations.                   |

### Example Usage

1. **Starting Simultaneous Audio Playback**
   To start audio playback between `alsa_output.pci-0000_0a_00.6.analog-stereo` and `bluez_output.D8_37_3B_09_3D_C0.1`, run:
   ```bash
   simuaudio -l
   ```
   look for what device is running and what device is suspended, then run:
   ```bash
   # simuaudio -s suspended_device_name or idle_device_name
   simuaudio -s alsa_output.pci-0000_0a_00.6.analog-stereo
   ```

   This will loop the audio from the capture device to the playback device, allowing both devices to play the same audio simultaneously.

2. **Killing Active Loopbacks**
   To stop all active `pw-loopback` processes, run:
   ```bash
   simuaudio -k
   ```

   This will kill any active `pw-loopback` processes that are running, stopping the audio playback between the devices.

3. **Listing Audio Devices**
   To list all available audio devices (sinks):
   ```bash
   simuaudio -l
   ```

   This will display all available audio output devices on your system.

### Additional Notes

- Devices must be specified correctly for the playback to work. Make sure the `node.target` names are valid and available in your system.
- The program checks if the device is already active in a loopback and avoids starting a new loopback if it's already running for that device.
- The `-s` flag allows you to specify multiple devices for simultaneous audio playback.
