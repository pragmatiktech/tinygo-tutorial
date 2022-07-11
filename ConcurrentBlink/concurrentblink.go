// ConcurrentBlink.go

package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	go printHello()

	for {
		led.High()
		time.Sleep(500 * time.Millisecond)
		led.Low()
		time.Sleep(500 * time.Millisecond)
	}
}

func printHello() {
	for {
		println("hello concurrently")
		time.Sleep(time.Second)
	}
}
