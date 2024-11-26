package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

func initKeyboardListener(stop chan os.Signal) error {
	if err := keyboard.Open(); err != nil {
		return err
	}

	go func() {
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				log.Printf("Error reading key: %v", err)
				close(stop)
				return
			}

			if key == keyboard.KeyEsc {
				close(stop)
				return
			}

			if char >= '1' && char <= '9' {
				mode := RunMode(char - '0')
				fmt.Printf("\nSwitched to: %s\n", modeDescriptions[mode])
				setMode(mode)
			}
		}
	}()
	return nil
}

func closeKeyboardListener() {
	keyboard.Close()
}
