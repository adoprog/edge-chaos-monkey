package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var mu sync.RWMutex // Mutex to safely access the current mode

func main() {
	// Handle clean shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Initialize keyboard input handling
	if err := initKeyboardListener(stop); err != nil {
		fmt.Printf("Failed to initialize keyboard listener: %v\n", err)
		return
	}
	defer closeKeyboardListener()

	// Print available modes
	printAvailableModes()

	// Start the proxy server
	go startProxyServer()

	// Main loop
	fmt.Printf("Current mode: %s\n", getModeDescription())
	fmt.Println("\nPress '1'-'9' to switch modes. Press 'Esc' to quit.")
	for {
		select {
		case <-stop:
			fmt.Println("Shutting down proxy server...")
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func printAvailableModes() {
	orderedModes := []RunMode{Mode1, Mode2, Mode3, Mode4, Mode5, Mode6, Mode7, Mode8, Mode9}
	for _, mode := range orderedModes {
		fmt.Printf("%d: %s\n", mode, modeDescriptions[mode])
	}
}
