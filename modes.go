package main

import "sync"

// RunMode represents the current run mode of the program.
type RunMode int

const (
	Mode1 RunMode = iota + 1
	Mode2
	Mode3
	Mode4
	Mode5
	Mode6
	Mode7
	Mode8
	Mode9
)

var modeDescriptions = map[RunMode]string{
	Mode1: "Mode 1: Basic Proxy Mode",
	Mode2: "Mode 2: Payload Logging Enabled (basic proxy, extra logging)",
	Mode3: "Mode 3: Request Throttling Mode (all requests return 429 error)",
	Mode4: "Mode 4: Slow Connection Mode (requests will take longer to process)",
	Mode5: "Mode 5: Server-side Error Mode (all requests return 500 error)",
	Mode6: "Mode 6: Chaos Mode (randomly return 500s, 429s, or slow responses)",
	Mode7: "Mode 7: Reserved for future use",
	Mode8: "Mode 8: Reserved for future use",
	Mode9: "Mode 9: Reserved for future use",
}

var (
	currentMode RunMode = Mode1
	modeMutex   sync.RWMutex
)

func setMode(mode RunMode) {
	modeMutex.Lock()
	defer modeMutex.Unlock()
	currentMode = mode
}

func getModeDescription() string {
	modeMutex.RLock()
	defer modeMutex.RUnlock()
	return modeDescriptions[currentMode]
}
