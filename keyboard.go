package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gvalkov/golang-evdev"
)

const (
	maxInputs = 50
)

func openDevice(deviceType ...string) *evdev.InputDevice {
	var device evdev.InputDevice
	for index := 0; index < maxInputs; index++ {
		device, err := evdev.Open(fmt.Sprintf(
			"/dev/input/event%d",
			index,
		))
		if err != nil || device == nil {
			return nil
		}
		for _, test := range deviceType {
			if strings.Contains(strings.ToLower(device.Name), test) {
				return device
			}
		}
	}
	return &device
}

func main() {
	keyboard := openDevice("keyboard")
	if keyboard == nil {
		log.Println("No keyboard detected")
	}
	mouse := openDevice("mouse", "touchpad")
	if mouse == nil {
		log.Println("No mouse detected")
	}
	fmt.Println(keyboard)
	fmt.Println(mouse)
	go func() {
		for {
			event, err := keyboard.ReadOne()
			fmt.Println("keyboard", event, err)
		}
	}()
	func() {
		for {
			event, err := mouse.ReadOne()
			fmt.Println("mouse", event, err)
		}
	}()
}
