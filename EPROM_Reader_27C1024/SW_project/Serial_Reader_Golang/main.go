/*   Read Arduino Serial and save bin format file
EPROM Reader 27C1024
Date: 2024_02_14   Gustavo Murta
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
		Name:        "COM15",
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

	s.Flush() // garbage clean at serial port

	fw, err := os.Create("EPROM_87F4794.bin") // create a bin format file up to 512K Bytes

	if err != nil { // if any error
		log.Fatal(err) // cancel and print error message
	}

	defer fw.Close() // close file at end

	buf := make([]byte, 1)    // define read buffer
	bufROM := make([]byte, 0) // define ROM buffer 16K x 8 bits

	log.Printf(" Read EPROM start")

	// 256K = 262144   512K=524288
	// 128K=131072     8K=8192   64K=65536     32K=32768  16K=16384

	for i := 0; i < (65536 * 2); i++ { // read 64K block with 16 bits data
		n, err := s.Read(buf) // read serial port data
		if err != nil {       // if any error
			log.Fatal(err, n) // cancel and print error message
		}
		bufROM = append(bufROM, buf...) // filling the ROM buffer
	}

	log.Printf(" EPROM Reading ends ") // print

	fmt.Printf("%s", hex.Dump(bufROM[:])) // print HEX dump with ASCII data

	for _, data := range bufROM {

		err := binary.Write(fw, binary.LittleEndian, data) // write ROM buffer into BIN format file

		if err != nil { // if any error
			log.Fatal(err) // cancel and print error message
		}
	}
	log.Printf(" Save file OK! ")

}
