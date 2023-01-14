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

	f, err := os.Open("basic-4000_commodore128.bin")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 16384)

	for {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		fmt.Printf("%s", hex.Dump(buf))
	}
}
