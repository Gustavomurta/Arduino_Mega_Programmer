/*   Read Serial and save bin format file
Date: 2023_02_05    Gustavo Murta
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
	} // define a config da porta serial - timeout de 1 seg

	s, err := serial.OpenPort(c) // abre a porta serial
	if err != nil {              // se apresentar algum erro
		log.Fatal(err) // cancela e imprime o erro
	}

	s.Flush() // limpa lixo na porta serial

	fw, err := os.Create("C128_ROM_BasicLOW.bin") // create a bin format file  16K x 8 bits

	if err != nil { // if any error
		log.Fatal(err) // print error message
	}

	defer fw.Close() // close file at end

	buf := make([]byte, 1)    // define o tamanho do buffer de leitura
	bufROM := make([]byte, 0) // create buffer ROM 16K x 8

	log.Printf(" Inicia a leitura da ROM")

	for i := 0; i < 16384; i++ { // efetua 16 vezes a leitura 16384 bytes
		n, err := s.Read(buf) // leitura dos dados na serial
		if err != nil {       // se apresentar algum erro
			log.Fatal(err, n) // imprime o erro
		}
		bufROM = append(bufROM, buf...) // filling the buffer
	}

	nw, err := s.Write([]byte("0")) // ascii 0 = 0x30
	if err != nil {
		log.Fatal(err, nw)
	}

	log.Printf(" Fim da leitura da ROM") // print

	fmt.Printf("%s", hex.Dump(bufROM[:])) // print HEX dump

	for _, data := range bufROM {

		err := binary.Write(fw, binary.LittleEndian, data) // write buffer into BIN format file

		if err != nil { // if any error
			log.Fatal(err) // print error message
		}
	}
	log.Printf(" Fim da gravação do arquivo")

}
