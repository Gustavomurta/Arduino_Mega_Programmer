/*   Go read binary file
     https://zetcode.com/golang/readfile/
*/

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("basic-4000_commodore128.bin") // open binary file

	if err != nil { // if any error
		log.Fatal(err) // print error message
	}

	defer f.Close() // close file at end

	reader := bufio.NewReader(f)
	buf := make([]byte, 16384) // create buffer  16384 bytes

	for {
		_, err := reader.Read(buf) // reaf file into buffer

		if err != nil { // if any error
			if err != io.EOF { // except end of file
				fmt.Println(err) // print error message
			}
			break
		}

		fmt.Printf("%s", hex.Dump(buf)) // print buffer in console
	}
}

