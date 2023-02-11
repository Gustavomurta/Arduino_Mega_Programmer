// ROM Reader ver1 - Arduino MEGA
// Gustavo Murta   2023/02/10
// Tinygo - version: 0.26.0

package main

import (
	"device/avr"
	"machine"
	"time"
)

var addressFF, address0F, addressF0 int
var delayPeriod = 1 * time.Microsecond // 1 us delay

func main() {

	machine.UART0.Configure(machine.UARTConfig{BaudRate: 115200}) // setup UART0 = 115200 Bps

	avr.DDRA.Set(0xFF) // configure PORTA pins as output - Memory Address A0 to A7
	avr.DDRC.Set(0xFF) // configure PORTC pins as output - Memory Address A8 to A15
	avr.DDRB.Set(0x0F) // configure PORT B0 to B3 pins as output
	avr.DDRG.Set(0x0F) // configure PORT G0 to G3 pins as output
	avr.DDRL.Set(0x00) // configure PORTL pins as input - Memory Data D0 to D7

	avr.PORTA.Set(0x00) // set Memory address A0 to A7 = 0
	avr.PORTC.Set(0xC0) // set Memory address A8 to A13 = 0 / A15 e A14 = 1
	avr.PORTB.Set(0x08) // set A16 to A18 = 0 and VPP = 1
	avr.PORTG.Set(0x0F) // set -CE(PG1) and -OE(PG0) = 1

	for addressFF = 0; addressFF < 16384; addressFF = addressFF + 256 { // block incremment 256 Bytes
		for addressF0 = 0; addressF0 < 256; addressF0 = addressF0 + 16 { // block incremment 16 Bytes
			for address0F = 0; address0F < 16; address0F++ { // addresses 0 to 15

				avr.PORTC.Set(uint8(addressFF>>8) | 0xC0)   // shift right 8 bits - Memory address A8 to A13 / bitwise OR A14 e A15 = 1
				avr.PORTA.Set(uint8(addressF0 + address0F)) // set Memory Address A0 to A7
				time.Sleep(delayPeriod)                     // delay 1 us
				avr.PORTG.Set(0x00)                         // set -CE(PG1) = 0 and -OE(PG0) = 0
				time.Sleep(delayPeriod)                     // delay 1 us

				epromData := (avr.PINL.Get())      // read Data Memory - D0 to D7
				machine.UART0.WriteByte(epromData) // send byte to serial port 0

				avr.PORTG.Set(0x01)     // set -CE(PG1) = 0 and -OE(PG0) = 1
				time.Sleep(delayPeriod) // delay 1 us
				avr.PORTA.Set(0x00)     // set Memory address A0 to A7 = 0
				avr.PORTC.Set(0xC0)     // set Memory Address A8 to A13 = 0 / A14 e A15 = 1
			}
		}
	}
	avr.PORTA.Set(0x00) // set Memory address A0 to A7 = 0
	avr.PORTC.Set(0x00) // set Memory Address A8 to A15 = 0
	avr.PORTG.Set(0x0F) // set -CE(PG1) and -OE(PG0) = 1
}
