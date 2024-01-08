package main

import (
	"fmt"
	"io" // new io package
	"strings"
)

func main() {

	// read end of a stream of data
	r := strings.NewReader("Hello, Reader!")

	// making a a byte slice of 8 bytes
	b := make([]byte, 8)
	
	// recall this is an infinite loop
	for {

		// read returns # of bytes read and an error value
		// it also populates our byte slice with our data

		// in other words, we are reading 8 bytes at a time
		n,err := r.Read(b)


		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])
		
		//if we get EOF, end loop
		if err == io.EOF {
			break
		}
	}
}