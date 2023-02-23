// EPROM 512K READER - Arduino MEGA
// Gustavo Murta   2023/02/22
// Tinygo - version: 0.26.0

package main

import (
	"device/avr"
	"machine"
	"time"
)

var address uint32
var block64K uint8
var delayPeriod = 1 * time.Microsecond // 1 us delay

func main() {

	machine.UART0.Configure(machine.UARTConfig{BaudRate: 115200}) // setup UART0 = 115200 Bps

	avr.DDRA.Set(0xFF) // configure PORTA pins as output - Memory Address pins A0 to A7
	avr.DDRC.Set(0xFF) // configure PORTC pins as output - Memory Address pins A8 to A15
	avr.DDRB.Set(0x0F) // configure PORT B0 to B3 pins as output - Memory Adress pins A16 to A18 and VPP
	avr.DDRG.Set(0x0F) // configure PORT G0 to G3 pins as output - CE and OE pins
	avr.DDRL.Set(0x00) // configure PORTL pins as input - Memory Data D0 to D7

	avr.PORTA.Set(0x00) // set Memory address A0 to A7 = 0
	avr.PORTC.Set(0x00) // set Memory address A8 to A15 = 0
	avr.PORTB.Set(0x08) // set A16 to A18 = 0 and VPP = 1
	avr.PORTG.Set(0x0F) // set -CE(PG1) and -OE(PG0) = 1

	for block64K = 0; block64K < 2; block64K++ { // count blocks of 64K => 128K
		avr.PORTB.Set(block64K | 0x08) // Memory address A16 to A18 / bitwise OR VPP = 1

		for address = 0; address < 65536; address++ { // read 64K block

			avr.PORTC.Set(uint8(address >> 8)) // shift right 8 bits - Memory address A8 to A15
			avr.PORTA.Set(uint8(address))      // set Memory Address A0 to A7
			time.Sleep(delayPeriod)            // delay 1 us
			avr.PORTG.Set(0x00)                // set -CE(PG1) = 0 and -OE(PG0) = 0
			time.Sleep(delayPeriod)            // delay 1 us

			epromData := (avr.PINL.Get())      // read Data Memory - D0 to D7
			machine.UART0.WriteByte(epromData) // send byte to serial port 0

			avr.PORTG.Set(0x01)     // set -CE(PG1) = 0 and -OE(PG0) = 1
			time.Sleep(delayPeriod) // delay 1 us
			avr.PORTA.Set(0x00)     // set Memory address A0 to A7 = 0
			avr.PORTC.Set(0x00)     // set Memory Address A8 to A15 = 0
		}
	}
	avr.PORTA.Set(0x00) // set Memory address A0 to A7 = 0
	avr.PORTC.Set(0x00) // set Memory Address A8 to A15 = 0
	avr.PORTB.Set(0x08) // set A16 to A18 = 0 and VPP = 1
	avr.PORTG.Set(0x0F) // set -CE(PG1) and -OE(PG0) = 1
}
