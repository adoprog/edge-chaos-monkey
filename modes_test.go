// FILE: modes_test.go
package main

import (
	"testing"
)

func TestSetMode(t *testing.T) {
	setMode(Mode1)
	modeMutex.RLock()
	defer modeMutex.RUnlock()
	if currentMode != Mode1 {
		t.Errorf("expected currentMode to be %v, got %v", Mode1, currentMode)
	}
}

func TestGetModeDescription(t *testing.T) {
	setMode(Mode1)
	expectedDescription := "Mode 1: Basic Proxy Mode"
	if description := getModeDescription(); description != expectedDescription {
		t.Errorf("expected description to be %v, got %v", expectedDescription, description)
	}
}
