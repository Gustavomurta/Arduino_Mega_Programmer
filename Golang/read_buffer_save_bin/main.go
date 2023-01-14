/*   Go read buffer (bin file) and save bin format file
     https://zetcode.com/golang/readfile/
     https://zetcode.com/golang/writefile/
     https://pkg.go.dev/encoding/binary#example-Write
*/

package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	fr, err := os.Open("87F4794.bin") // open file

	if err != nil {
		log.Fatal(err)
	}

	defer fr.Close() // close file at end

	reader := bufio.NewReader(fr)
	buf := make([]byte, 131072) // create buffer  16384 131072
	fmt.Println()               // print line

	for {
		_, err := reader.Read(buf) // reaf file into buffer

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		fmt.Printf("%s", hex.Dump(buf)) // print HEX dump
	}

	fw, err := os.Create("rom_87F4794.bin") // create a file

	if err != nil { // if any error
		log.Fatal(err) // print error message
	}

	defer fw.Close() // close file at end

	for _, data := range buf {

		err := binary.Write(fw, binary.LittleEndian, data) // write buffer into BIN format file

		if err != nil { // if any error
			log.Fatal(err) // print error message
		}
	}
	fmt.Println("done") // print
}
