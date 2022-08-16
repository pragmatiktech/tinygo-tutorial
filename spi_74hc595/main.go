package main

import (
    "machine"
    "time"
)

const (
    rclk = machine.GP20
)

var sr = machine.SPI0

func write(val int) {
    rclk.Low()
    sr.Tx([]byte{byte(val)}, nil)
    rclk.High()
    rclk.Low()
}

func configure() {
    sr.Configure(machine.SPIConfig{
        Frequency: 100000,
        LSBFirst:  false,
        Mode:      1,
        DataBits:  8,
        SCK:       machine.GP18,
        SDO:       machine.GP19,
    })
    rclk.Configure(machine.PinConfig{
        Mode: machine.PinOutput,
    })
    rclk.Low()
}

func main() {
    configure()
    for i := 0; i < 256; i++ {
        write(i)
        time.Sleep(100 * time.Millisecond)
    }
    time.Sleep(time.Second)
    write(0)
}
