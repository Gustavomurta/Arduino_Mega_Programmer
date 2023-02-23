/*   Read Arduino Serial and save bin format file
EPROM READER 512K
Date: 2023_02_22    Gustavo Murta
https://zetcode.com/golang/readfile/
https://zetcode.com/golang/writefile/
https://pkg.go.dev/encoding/binary#example-Write

*/

package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tarm/serial"
)

func main() {

	c := &serial.Config{
		Name:        "COM11",
		Baud:        115200,
		ReadTimeout: time.Second * 1,
		Size:        8,
		Parity:      0,
		StopBits:    1,
	} // Win10 Serial port setup 115200 Bps- timeout 1 second

	s, err := serial.OpenPort(c) // opens serial port
	if err != nil {              // if any error
		log.Fatal(err) // cancel and print error message
	}

	s.Flush() // gargage clean at serial port

	fw, err := os.Create("EPROM_128K_ModemUSR.bin") // create a bin format file  128K x 8 bits

	if err != nil { // if any error
		log.Fatal(err) // cancel and print error message
	}

	defer fw.Close() // close file at end

	buf := make([]byte, 1)    // define read buffer
	bufROM := make([]byte, 0) // define ROM buffer 16K x 8 bits

	log.Printf(" Read EPROM start")

	for i := 0; i < 131072; i++ { // read 512K bytes  => 128K
		n, err := s.Read(buf) // read serial port data
		if err != nil {       // if any error
			log.Fatal(err, n) // cancel and print error message
		}
		bufROM = append(bufROM, buf...) // filling the ROM buffer
	}

	log.Printf(" EPROM Reading ends ") // print

	fmt.Printf("%s", hex.Dump(bufROM[:])) // print HEX dump with ASCII

	for _, data := range bufROM {

		err := binary.Write(fw, binary.LittleEndian, data) // write ROM buffer into BIN format file

		if err != nil { // if any error
			log.Fatal(err) // cancel and print error message
		}
	}
	log.Printf(" Save file OK! ")

}
