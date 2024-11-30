package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Declare variables
	containerName := "lightweight-webserver"
	imageName := "httpd:alpine"
	networkName := "custom-bridge-network"
	subnet := "192.168.1.0/24"
	staticIP := "192.168.1.150"
	portMapping := "8085:85"

	// Helper to run commands
	runCommand := func(name string, args ...string) error {
		cmd := exec.Command(name, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		fmt.Printf("Running: %s %v\n", name, args)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		return err
	}

	// Step 1: Ensure Docker is installed
	fmt.Println("Checking if Docker is installed...")
	if err := runCommand("docker", "--version"); err != nil {
		fmt.Println("Error: Docker is not installed or not accessible.")
		os.Exit(1)
	}

	// Step 2: Create a custom Docker network (if it doesn't exist)
	fmt.Printf("Creating custom Docker network: %s with subnet: %s...\n", networkName, subnet)
	if err := runCommand("docker", "network", "create", "--subnet", subnet, networkName); err != nil {
		fmt.Println("Warning: Network might already exist or could not be created.")
	}

	// Step 3: Pull the lightweight web server image
	fmt.Println("Pulling the lightweight httpd:alpine image...")
	if err := runCommand("docker", "pull", imageName); err != nil {
		fmt.Println("Error pulling the image. Exiting.")
		os.Exit(1)
	}

	// Step 4: Stop any existing container
	fmt.Printf("Stopping the container: %s...\n", containerName)
	runCommand("docker", "stop", containerName)

	// Step 5: Remove any existing container
	fmt.Printf("Removing the container: %s...\n", containerName)
	runCommand("docker", "rm", containerName)

	// Step 6: Create the container with static IP and port mapping
	fmt.Printf("Creating a new container: %s with image: %s and static IP: %s...\n", containerName, imageName, staticIP)
	if err := runCommand(
		"docker", "create",
		"--network", networkName,
		"--ip", staticIP,
		"-p", portMapping,
		"--name", containerName,
		imageName,
	); err != nil {
		fmt.Println("Error creating the container. Exiting.")
		os.Exit(1)
	}

	// Step 7: Start the created container
	fmt.Printf("Starting the container: %s...\n", containerName)
	if err := runCommand("docker", "start", containerName); err != nil {
		fmt.Println("Error starting the container. Exiting.")
		os.Exit(1)
	}

	// Success message
	fmt.Printf("Lightweight web server deployed successfully with static IP: %s\n", staticIP)
}
