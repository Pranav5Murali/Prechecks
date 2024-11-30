package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Declare variables
	var containerName string = "lightweight-webserver"
	var imageName string = "httpd:alpine"
	var networkName string = "custom-bridge-network"
	var subnet string = "192.168.1.0/24"
	var staticIP string = "192.168.1.150"
	var portMapping string = "8085:85"
	var err error
	var cmd *exec.Cmd

	// Step 1: Create a custom Docker network (if it doesn't exist)
	fmt.Printf("Creating custom Docker network: %s with subnet: %s...\n", networkName, subnet)
	cmd = exec.Command("docker", "network", "create", "--subnet", subnet, networkName)
	cmd.Stdout = exec.Stdout
	cmd.Stderr = exec.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Warning: Network might already exist. Error:", err)
	}

	// Step 2: Pull the lightweight web server image
	fmt.Println("Pulling the lightweight httpd:alpine image...")
	cmd = exec.Command("docker", "pull", imageName)
	cmd.Stdout = exec.Stdout
	cmd.Stderr = exec.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error pulling the image:", err)
		return
	}

	// Step 3: Stop any existing container (ignore error if it doesn't exist)
	fmt.Printf("Stopping the container: %s...\n", containerName)
	cmd = exec.Command("docker", "stop", containerName)
	cmd.Stdout = exec.Stdout
	cmd.Stderr = exec.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Warning: Error stopping container (might not exist):", err)
	}

	// Step 4: Remove any existing container (ignore error if it doesn't exist)
	fmt.Printf("Removing the container: %s...\n", containerName)
	cmd = exec.Command("docker", "rm", containerName)
	cmd.Stdout = exec.Stdout
	cmd.Stderr = exec.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Warning: Error removing container (might not exist):", err)
	}

	// Step 5a: Create the container with static IP and port mapping
	fmt.Printf("Creating a new container: %s with image: %s and static IP: %s...\n", containerName, imageName, staticIP)
	cmd = exec.Command(
		"docker", "create",
		"--network", networkName,
		"--ip", staticIP,
		"-p", portMapping,
		"--name", containerName,
		imageName,
	)
	cmd.Stdout = exec.Stdout
	cmd.Stderr = exec.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error creating the container:", err)
		return
	}

	// Step 5b: Start the created container
	fmt.Printf("Starting the container: %s...\n", containerName)
	cmd = exec.Command("docker", "start", containerName)
	cmd.Stdout = exec.Stdout
	cmd.Stderr = exec.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error starting the container:", err)
		return
	}

	// Success message
	fmt.Printf("Lightweight web server deployed successfully with static IP: %s\n", staticIP)
}
