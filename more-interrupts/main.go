package main

import (
	"machine"
	"time"
)

const (
	led  = machine.GP16
	led2 = machine.GP15
	btn  = machine.GP17
)

// This function configures the GPIO peripherals, two LEDs as output, and
// the push button as an input with a pullup resistor using the internal pullup.
func configure() {
	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	led2.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	btn.Configure(machine.PinConfig{
		Mode: machine.PinInputPullup,
	})
}

// isr is the function that "services" the interrupt,
// or is the "callback" function when an interrupt is fired on the pin connected
// to the push button. In this case, the function merely sets the LED to the
// inverse of the push button state.
func isr(p machine.Pin) {
	// disable the interrupt service routine since that's a good thing to do.
	btn.SetInterrupt(machine.PinFalling|machine.PinRising, nil)

	led.Set(!p.Get()) // this hasn't changed from our original program.

	// simulate a long-running process
	for i := 0; i < 10000; i++ {
		print(i)
	}

	// re-enable the interrupt service routine.
	btn.SetInterrupt(machine.PinFalling|machine.PinRising, isr)
}

func main() {
	// Configure the pins.
	configure()

	// Register the interrupt service routine (ISR)
	btn.SetInterrupt(machine.PinFalling|machine.PinRising, isr)

	// Finally, just loop indefinitely while blinking the green LED.
	for {
		led2.High()
		time.Sleep(500 * time.Millisecond)
		led2.Low()
		time.Sleep(500 * time.Millisecond)
	}
}
