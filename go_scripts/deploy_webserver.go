package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Declare configuration variables
	var containerName string = "lightweight-webserver"
	var imageName string = "httpd:alpine"
	var networkName string = "custom-bridge-network"
	var subnet string = "192.168.1.0/24"
	var staticIP string = "192.168.1.150"
	var portMapping string = "8085:85"


	// Step 1: Check if the Docker network exists, create if not
	fmt.Printf("Ensuring Docker network: %s with subnet: %s exists...\n", networkName, subnet)
	var cmdInspectNetwork *exec.Cmd = exec.Command("docker", "network", "inspect", networkName)
	cmdInspectNetwork.Stdout = os.Stdout
	cmdInspectNetwork.Stderr = os.Stderr
	var errInspectNetwork error = cmdInspectNetwork.Run()
	if errInspectNetwork != nil {
		fmt.Printf("Network %s does not exist. Creating...\n", networkName)
		var cmdCreateNetwork *exec.Cmd = exec.Command("docker", "network", "create", "--subnet", subnet, networkName)
		cmdCreateNetwork.Stdout = os.Stdout
		cmdCreateNetwork.Stderr = os.Stderr
		var errCreateNetwork error = cmdCreateNetwork.Run()
		if errCreateNetwork != nil {
			fmt.Println("Error: Failed to create Docker network. Exiting.")
			os.Exit(1)
		}
	} else {
		fmt.Println("Network already exists. Skipping creation.")
	}

	// Step 2: Pull the lightweight web server image
	fmt.Println("Pulling the lightweight httpd:alpine image...")
	var cmdPullImage *exec.Cmd = exec.Command("docker", "pull", imageName)
	cmdPullImage.Stdout = os.Stdout
	cmdPullImage.Stderr = os.Stderr
	var errPullImage error = cmdPullImage.Run()
	if errPullImage != nil {
		fmt.Println("Error: Failed to pull the Docker image. Exiting.")
		os.Exit(1)
	}

	// Step 3: Create a new container with a static IP and port mapping
	fmt.Printf("Creating a new container: %s with image: %s and static IP: %s...\n", containerName, imageName, staticIP)
	var cmdCreateContainer *exec.Cmd = exec.Command(
		"docker", "create",
		"--network", networkName,
		"--ip", staticIP,
		"-p", portMapping,
		"--name", containerName,
		imageName,
	)
	cmdCreateContainer.Stdout = os.Stdout
	cmdCreateContainer.Stderr = os.Stderr
	var errCreateContainer error = cmdCreateContainer.Run()
	if errCreateContainer != nil {
		fmt.Println("Error: Failed to create the container. Exiting.")
		os.Exit(1)
	}

	// Step 4: Start the newly created container
	fmt.Printf("Starting the container: %s...\n", containerName)
	var cmdStartContainer *exec.Cmd = exec.Command("docker", "start", containerName)
	cmdStartContainer.Stdout = os.Stdout
	cmdStartContainer.Stderr = os.Stderr
	var errStartContainer error = cmdStartContainer.Run()
	if errStartContainer != nil {
		fmt.Println("Error: Failed to start the container. Exiting.")
		os.Exit(1)
	}

	// Success message
	fmt.Printf("Container %s deployed successfully with static IP: %s\n", containerName, staticIP)
}
