/*   Go read binary file
     https://zetcode.com/golang/readfile/
	 https://zetcode.com/golang/writefile/
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

	fr, err := os.Open("basic-4000_commodore128.bin") // open file

	if err != nil {
		log.Fatal(err)
	}

	defer fr.Close() // close file at end

	reader := bufio.NewReader(fr)
	buf := make([]byte, 16384) // create buffer

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

	fw, err := os.Create("rom.txt") // create a file

	if err != nil { // if any error
		log.Fatal(err) // print error message
	}

	defer fw.Close() // close file at end

	n, err := fmt.Fprintf(fw, "%s", hex.Dump(buf)) // write buffer into TXT format file

	if err != nil { // if any error
		log.Fatal(err) // print error message
	}

	fmt.Println(n, "bytes written") // print
	fmt.Println("done")             // print
}
