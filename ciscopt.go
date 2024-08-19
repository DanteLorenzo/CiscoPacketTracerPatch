package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/DanteLorenzo/CiscoPacketTracerPatch/app/logs"      // Importing the logs package
	"github.com/DanteLorenzo/CiscoPacketTracerPatch/app/processor" // Importing the processor package
	"github.com/DanteLorenzo/CiscoPacketTracerPatch/app/types"     // Importing the types package
	"github.com/fatih/color"                      // Importing the color package for colored output
)

func main() {
	// Record the start time of the patching process
	start := time.Now()

	// Define color styles for logging
	info := color.New(color.FgCyan)
	success := color.New(color.FgGreen)
	errorColor := color.New(color.FgRed)
	startEndColor := color.New(color.FgGreen)
	durationColor := color.New(color.FgGreen)

	// Log the start of the patching process with a timestamp
	logs.LogInfo("Starting the patching process at "+start.Format(time.RFC1123), startEndColor)

	// Get the directory of the running program
	execPath, err := os.Executable()
	if err != nil {
		logs.LogFatal("Failed to get the executable path: %v", err)
	}
	dirPath := filepath.Dir(execPath)

	// Search for the "PacketTracer.exe" file within the directory
	var filePath string
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "PacketTracer.exe" {
			filePath = path
			return filepath.SkipDir // Stop walking the directory once the file is found
		}
		return nil
	})
	if err != nil {
		logs.LogFatal("Error walking the directory: %v", err) // Log fatal error if walking the directory fails
	}

	if filePath == "" {
		logs.LogFatal("PacketTracer.exe not found in the program's directory") // Log fatal error if the file is not found
	}

	// Log the found file path
	logs.LogInfo("Found PacketTracer.exe at: "+filePath, success)

	// Define the patterns directly in the code using types.Pattern
	patterns := []types.Pattern{
		{
			OldPattern: "90 49 8B CF E8 ?? ?? ?? ?? 84 C0 0F 85 ?? ?? ?? ?? 49 8B CF E8 ?? ?? ?? ?? 84 C0 0F 85 ?? ?? ?? ??",
			NewPattern: "90 49 8B CF E8 ?? ?? ?? ?? 84 C0 0F 85 ?? ?? ?? ?? 49 8B CF E8 ?? ?? ?? ?? 84 C0 0F 84 ?? ?? ?? ??",
		},
		{
			OldPattern: "49 8B CF E8 ?? ?? ?? ?? 84 C0 0F 85 ?? ?? ?? ?? BA 02 00 00 00",
			NewPattern: "49 8B CF E8 ?? ?? ?? ?? 84 C0 0F 84 ?? ?? ?? ?? BA 02 00 00 00",
		},
	}

	dir, fileName := filepath.Split(filePath)            // Split the file path into directory and file name
	originFileName := fileName + ".origin"               // Create the name for the backup of the original file
	originFilePath := filepath.Join(dir, originFileName) // Combine directory and new file name for the backup
	newFilePath := filePath                              // The new file path remains the same as the original

	// Rename the original file to keep a backup
	err = os.Rename(filePath, originFilePath)
	if err != nil {
		logs.LogFatal("Failed to rename original file: %v", err) // Log fatal error if renaming fails
	}

	// Read the contents of the original (now renamed) file
	fileData, err := os.ReadFile(originFilePath)
	if err != nil {
		logs.LogFatal("Failed to read file data: %v", err) // Log fatal error if reading the file fails
	}

	// Log the patterns found in the original file
	logs.LogInfo("Original file patterns:", info)
	for _, pattern := range patterns {
		// Print the patterns found in the original file
		processor.PrintPatterns(fileData, pattern.OldPattern, info, success, errorColor)
	}

	// Modify the file data by replacing the old patterns with new patterns
	modifiedData := processor.ModifyPatterns(fileData, patterns)

	// Create a new file with the modified data
	newFile, err := os.Create(newFilePath)
	if err != nil {
		logs.LogFatal("Failed to create new file: %v", err) // Log fatal error if creating the new file fails
	}
	defer newFile.Close() // Ensure the file is closed after writing is done

	// Write the modified data to the new file
	_, err = newFile.Write(modifiedData)
	if err != nil {
		logs.LogFatal("Failed to write modified data to new file: %v", err) // Log fatal error if writing fails
	}

	// Sync the new file to ensure all data is written to disk
	if err := newFile.Sync(); err != nil {
		logs.LogFatal("Error syncing file: %v", err) // Log fatal error if syncing fails
	}

	// Read the contents of the newly created file to verify modifications
	newFileData, err := os.ReadFile(newFilePath)
	if err != nil {
		logs.LogFatal("Failed to read modified file data: %v", err) // Log fatal error if reading the modified file fails
	}

	// Log the patterns found in the patched file
	logs.LogInfo("Patched file patterns:", info)
	for _, pattern := range patterns {
		// Print the patterns found in the patched file
		processor.PrintPatterns(newFileData, pattern.NewPattern, info, success, errorColor)
	}

	// Record the end time of the patching process
	end := time.Now()

	// Log the end of the patching process with a timestamp
	logs.LogInfo("Finished the patching process at "+end.Format(time.RFC1123), startEndColor)

	// Log the total duration of the patching process
	logs.LogInfo("Total duration: "+end.Sub(start).String(), durationColor)
}
