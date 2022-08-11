package main

import (
    "machine"
    "math/rand"
    "time"
)

const (
    yellowLed = machine.GP16
    greenLed  = machine.GP15
)

func configure() {
    yellowLed.Configure(machine.PinConfig{Mode: machine.PinOutput})
    greenLed.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

// Blinks an LED twice a second for a specified number of times.
func blink(p machine.Pin, n int) {
    for i := 0; i < n; i++ {
        p.High()
        time.Sleep(250 * time.Millisecond)
        p.Low()
        time.Sleep(250 * time.Millisecond)
    }
}

// This goroutine blinks the green LED `numBlinks` times, where
// numBlinks is specified by the message it receives via the channel.
func blinkGreen(c chan int) {
    for {
        // block until you receive a value via the channel
        numBlinks := <-c
        
        // then, blink the green LED that many times
        blink(greenLed, numBlinks)
        
        // return a random value to the `main` goroutine so that it blinks
        // the yellow LED as many times as this random value.
        numBlinks = rand.Intn(5-1) + 1
        c <- numBlinks
    }
}

func main() {
    // Create an unbuffered channel of ints.
    c := make(chan int)
    
    // Initialize the random seed
    rand.Seed(time.Now().UnixNano())
    
    // Generate a random number between 1 and 5
    n := rand.Intn(5-1) + 1
    
    configure()
    
    // Start the blinkGreen goroutine
    go blinkGreen(c)
    
    for {
        // First blink the yellow LED `n` times
        blink(yellowLed, n)
        
        // then, generate a random number between 1 and 5
        n = rand.Intn(5-1) + 1
        
        // send that random number via the channel to `blinkGreen`
        // so that it blinks the green LED that many times
        c <- n
        
        // wait for `blinkGreen` to return a random number to you
        n = <-c
    } // and repeat indefinitely
}
