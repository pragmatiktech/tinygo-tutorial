package main

import (
    "machine"
    "time"
)

func main() {
    // Initialize the white LED on GPIO16
    led := machine.GP16
    led.Configure(machine.PinConfig{
        Mode: machine.PinOutput,
    })
    
    // Initialize the green LED on GPIO15
    led2 := machine.GP15
    led2.Configure(machine.PinConfig{
        Mode: machine.PinOutput,
    })
    
    // Initialize the button on GPIO17 and set its internal pull-up resistor.
    btn := machine.GP17
    btn.Configure(machine.PinConfig{
        Mode: machine.PinInputPullup,
    })
    
    // Register the interrupt service routine (ISR) to set the status of the
    // white LED to the *inverse* of the status of the pushbutton pin.
    btn.SetInterrupt(machine.PinFalling|machine.PinRising,
        func(p machine.Pin) {
            led.Set(!p.Get())
        })
    
    // Finally, just loop indefinitely while blinking the green LED.
    for {
        led2.High()
        time.Sleep(500 * time.Millisecond)
        led2.Low()
        time.Sleep(500 * time.Millisecond)
    }
}
