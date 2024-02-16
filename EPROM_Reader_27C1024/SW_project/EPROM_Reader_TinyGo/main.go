// EPROM Reader 27C1024 - Arduino MEGA versão 2
// Gustavo Murta   2024_02_14
// Tinygo - version: 0.30.0
// tinygo flash -target=arduino-mega2560 main.go

package main

import (
	"device/avr"
	"machine"
	"time"
)

var address uint32
var delayPeriod = 1 * time.Microsecond // 1 us delay

func main() {

	machine.UART0.Configure(machine.UARTConfig{BaudRate: 115200}) // setup UART0 = 115200 Bps

	avr.DDRA.Set(0xFF) // configure PORTA pins as output - address lines A0 to A7
	avr.DDRC.Set(0xFF) // configure PORTC pins as output - address lines A8 to A15
	avr.DDRG.Set(0x0F) // configure PORTG pins 3,2,1 and 0 as output
	avr.DDRL.Set(0x00) // configure PORTL pins as input - data lines DQ0 to DQ7
	avr.DDRK.Set(0x00) // configure PORTL pins as input - data lines DQ8 to DQ15

	avr.PORTA.Set(0x00) // set Memory address A0 to A7 = 0
	avr.PORTC.Set(0x00) // set Memory address A8 to A15 = 0
	avr.PORTG.Set(0x03) // set -CE(PG1) and -OE(PG0) = 1

	// 256K = 262144   512K=524288
	// 128K=131072     8K=8192   64K=65536     32K=32768  16K=16384

	for address = 0; address < 65536; address++ { // read 64K block with 16 bits data

		avr.PORTC.Set(uint8(address >> 8)) // shift right 8 bits - Set Memory address A8 to A15
		avr.PORTA.Set(uint8(address))      // set Memory Address A0 to A7
		time.Sleep(delayPeriod)            // delay 1 us
		avr.PORTG.Set(0x00)                // set -CE(PG1) = 0 and -OE(PG0) = 0
		time.Sleep(delayPeriod)            // delay 1 us

		epromDataLow := (avr.PINL.Get())       // read Data Memory - DQ0 to DQ7
		epromDataHigh := (avr.PINK.Get())      // read Data Memory - DQ8 to DQ15
		machine.UART0.WriteByte(epromDataLow)  // send Low byte to serial port 0
		machine.UART0.WriteByte(epromDataHigh) // send High byte to serial port 0

		avr.PORTG.Set(0x01)     // set -CE(PG1) = 0 and -OE(PG0) = 1
		time.Sleep(delayPeriod) // delay 1 us
		avr.PORTA.Set(0x00)     // set Memory address A0 to A7 = 0
		avr.PORTC.Set(0x00)     // set Memory Address A8 to A15 = 0
	}

	avr.PORTA.Set(0x00) // set Memory address A0 to A7 = 0
	avr.PORTC.Set(0x00) // set Memory Address A8 to A15 = 0
	avr.PORTG.Set(0x03) // set -CE(PG1) and -OE(PG0) = 1
}
