# Cisco Packet Tracer Patcher

This Go application automates the process of patching the `PacketTracer.exe` file for Cisco Packet Tracer by locating and modifying specific byte patterns in the executable.

## Features

- **Automatic File Discovery**: The application automatically finds the `PacketTracer.exe` file in the directory where the program is executed.
- **Pattern Matching and Replacement**: Replaces specific byte patterns in the `PacketTracer.exe` file with new ones as defined in the code.
- **Backup Creation**: Automatically creates a backup of the original `PacketTracer.exe` file before applying any modifications.
- **Detailed Logging**: Logs the entire process with timestamps and colored output for easy debugging and monitoring.
- **Error Handling**: Provides robust error handling to ensure that any issues are logged and the program exits gracefully.

## Requirements

- **Windows Only**: This application is designed to work exclusively with the Windows version of Cisco Packet Tracer.
- **Administrator Privileges**: The application must be run with administrator privileges to ensure it can modify the `PacketTracer.exe` file.

## Supported Versions

The patcher is designed to work with the following versions of Cisco Packet Tracer:
- **8.2.1**
- **8.2.2**

## Usage

1. Download the latest release executable from the [Releases](https://github.com/DanteLorenzo/CiscoPacketTracerPatch/releases) page.

2. Place the downloaded `CiscoPacketTracerPatcher.exe` in the directory where `PacketTracer.exe` is located.

3. **Run the patcher with administrator privileges**:
    - Right-click the executable and select "Run as administrator."
    - Alternatively, run it from an elevated command prompt.

    ```sh
    ./CiscoPacketTracerPatcher.exe
    ```

The patcher will automatically locate the `PacketTracer.exe` file in its directory, create a backup, apply the necessary patches, and log the entire process.

## Logs

Logs are automatically generated and provide detailed information about the patching process, including:
- Start and end times
- Found patterns in the original and modified files
- File paths and error messages (if any)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
